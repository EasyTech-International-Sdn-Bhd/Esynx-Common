package client

import (
	"errors"
	"github.com/easytech-international-sdn-bhd/esynx-common/rest"
	rq "github.com/easytech-international-sdn-bhd/esynx-common/rest/request"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
)

type TokenData struct {
	AccessToken  string   `mapstructure:"AccessToken"`
	RefreshToken string   `mapstructure:"RefreshToken"`
	Roles        []string `mapstructure:"Roles"`
	Permissions  []string `mapstructure:"Permissions"`
}

type LoginResponse struct {
	Data    TokenData `json:"Data"`
	Message string    `json:"Message"`
	Status  bool      `json:"Status"`
}

type RestClientParams struct {
	Url      string
	AppName  string
	Username string
	Password string
}

type ApiClient struct {
	Reqwest *resty.Client
}

func NewApiClient(param *RestClientParams) (*ApiClient, error) {
	const authHeaderName = "X-App-Name"
	const acceptHeader = "Accept"
	const acceptHeaderValue = "application/json"
	const authSchema = "Bearer"

	client := resty.New()
	client.
		SetBaseURL(param.Url).
		SetDebug(false).
		SetRetryCount(1).
		SetHeader(authHeaderName, param.AppName).
		SetHeader(acceptHeader, acceptHeaderValue).
		SetContentLength(true).
		EnableTrace().
		OnError(func(req *resty.Request, err error) {
			log.Println(err)
		})

	response, err := loginUser(client, param)
	if err != nil {
		return nil, err
	}

	client.
		SetAuthScheme(authSchema).
		SetAuthToken(response.Data.AccessToken).
		OnBeforeRequest(handleAuth(param, response)).
		AddRetryCondition(func(r *resty.Response, err error) bool {
			return err != nil || r.StatusCode() == http.StatusUnauthorized
		})

	return &ApiClient{Reqwest: client}, nil
}

func loginUser(client *resty.Client, param *RestClientParams) (*LoginResponse, error) {
	var response *LoginResponse
	r, err := client.
		R().
		SetBody(map[string]string{"username": param.Username, "password": param.Password}).
		SetResult(&response).
		Post(rest.AuthLogin)
	if err != nil {
		return nil, err
	}
	if r.IsError() {
		return nil, errors.New(r.String())
	}
	return response, nil
}

func handleAuth(param *RestClientParams, response *LoginResponse) func(client *resty.Client, request *resty.Request) error {
	return func(client *resty.Client, request *resty.Request) error {
		var authResp *ReferenceBasedResponse
		reqAcs := rq.ClientAccessTokenForm{AccessToken: response.Data.AccessToken}
		acsResp, err := request.SetBody(reqAcs).SetResult(&authResp).Post(rest.AuthCheckAccessToken)
		if err != nil || acsResp.IsError() {
			reqRfs := rq.ClientRefreshTokenForm{RefreshToken: response.Data.RefreshToken}
			rfsResp, err := request.SetBody(reqRfs).SetResult(&authResp).Post(rest.AuthCheckRefreshToken)
			if err != nil || rfsResp.IsError() {
				response, err = loginUser(client, param)
				if err != nil {
					return err
				}
				request.SetAuthToken(response.Data.AccessToken)
			}
		}
		return nil
	}
}
