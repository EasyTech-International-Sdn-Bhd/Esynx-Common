package client

import (
	"errors"
	"github.com/easytech-international-sdn-bhd/esynx-common/rest"
	rq "github.com/easytech-international-sdn-bhd/esynx-common/rest/request"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
	"strings"
)

// TokenData holds the token details
type TokenData struct {
	AccessToken  string   `json:"accessToken"`
	RefreshToken string   `json:"refreshToken"`
	Roles        []string `json:"roles"`
	Permissions  []string `json:"permissions"`
	ClientId     string   `json:"clientId"`
}

// RefreshTokenData holds the refresh token details
type RefreshTokenData struct {
	AccessToken string `json:"accessToken"`
}

// RefreshTokenResponse holds the response structure for refresh token
type RefreshTokenResponse struct {
	Data    RefreshTokenData `json:"Data"`
	Message string           `json:"Message"`
	Status  bool             `json:"Status"`
}

// LoginResponse holds the response structure for login
type LoginResponse struct {
	Data    TokenData `json:"Data"`
	Message string    `json:"Message"`
	Status  bool      `json:"Status"`
}

// RestClientParams holds the parameters for the REST client
type RestClientParams struct {
	URL      string `json:"url"`
	AppName  string `json:"appName"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// ApiClient holds the REST client
type ApiClient struct {
	Reqwest  *resty.Client
	ClientId string
}

// NewApiClient creates a new API client
func NewApiClient(params *RestClientParams, debug ...bool) (*ApiClient, error) {
	const (
		authHeaderName    = "X-App-Name"
		acceptHeader      = "Accept"
		acceptHeaderValue = "application/json"
		authSchema        = "Bearer"
	)
	ignore := []string{rest.AuthCheckAccessToken, rest.AuthCheckRefreshToken, rest.AuthRefreshToken, rest.AuthLogin}

	client := resty.New().
		SetBaseURL(params.URL).
		SetDebug(len(debug) > 0).
		SetRetryCount(1).
		SetHeader(authHeaderName, params.AppName).
		SetHeader(acceptHeader, acceptHeaderValue).
		SetContentLength(true).
		EnableTrace().
		OnError(func(req *resty.Request, err error) {
			log.Println(err)
		}).
		AddRetryCondition(func(response *resty.Response, err error) bool {
			for _, s := range ignore {
				if strings.Contains(response.Request.URL, s) {
					return false
				}
			}
			return err != nil || response.StatusCode() == http.StatusUnauthorized
		})

	loginResponse, err := loginUser(client, params)
	if err != nil {
		return nil, err
	}

	client.SetAuthScheme(authSchema).
		SetAuthToken(loginResponse.Data.AccessToken).
		OnBeforeRequest(func(clx *resty.Client, r *resty.Request) error {
			for _, s := range ignore {
				if strings.Contains(r.URL, s) {
					return nil
				}
			}
			return handleAuth(params, loginResponse)(clx, r)
		})

	return &ApiClient{Reqwest: client, ClientId: loginResponse.Data.ClientId}, nil
}

func loginUser(client *resty.Client, params *RestClientParams) (*LoginResponse, error) {
	var response *LoginResponse
	resp, err := client.R().
		SetBody(map[string]string{"username": params.Username, "password": params.Password}).
		SetResult(&response).
		Post(rest.AuthLogin)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, errors.New(resp.String())
	}

	return response, nil
}

func handleAuth(params *RestClientParams, loginResponse *LoginResponse) func(client *resty.Client, request *resty.Request) error {
	return func(client *resty.Client, request *resty.Request) error {
		var authResponse *ReferenceBasedResponse

		// Check if access token is still valid
		accessTokenForm := rq.ClientAccessTokenForm{AccessToken: loginResponse.Data.AccessToken}
		acsResp, err := client.R().
			SetBody(accessTokenForm).
			SetResult(&authResponse).
			Post(rest.AuthCheckAccessToken)

		if err == nil && acsResp.StatusCode() == http.StatusOK && authResponse.Status {
			// Access token is valid, proceed with the request
			return nil
		}

		// Access token is invalid; try refreshing it
		refreshTokenForm := rq.ClientRefreshTokenForm{RefreshToken: loginResponse.Data.RefreshToken}
		rfsResp, err := client.R().
			SetBody(refreshTokenForm).
			SetResult(&authResponse).
			Post(rest.AuthCheckRefreshToken)

		if err == nil && rfsResp.StatusCode() == http.StatusOK && authResponse.Status {
			var refreshResponse *RefreshTokenResponse

			rfsResp, err = client.R().
				SetBody(refreshTokenForm).
				SetResult(&refreshResponse).
				Post(rest.AuthRefreshToken)

			if err == nil && rfsResp.StatusCode() == http.StatusOK && refreshResponse.Status {
				// Refresh token succeeded, update the token and proceed with the request
				loginResponse.Data.AccessToken = refreshResponse.Data.AccessToken
				request.SetAuthToken(loginResponse.Data.AccessToken)
				client.SetAuthToken(loginResponse.Data.AccessToken)
				return nil
			}
		}

		// Refresh token also expired, perform login to get new tokens
		loginResponse, err = loginUser(client, params)
		if err != nil {
			return err
		}

		// Set the new access token for the request
		request.SetAuthToken(loginResponse.Data.AccessToken)
		client.SetAuthToken(loginResponse.Data.AccessToken)
		return nil
	}
}
