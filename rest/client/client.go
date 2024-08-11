package client

import "github.com/go-resty/resty/v2"

type ApiClient struct {
	Reqwest *resty.Client
}

func NewApiClient(rq *resty.Client) *ApiClient {
	return &ApiClient{Reqwest: rq}
}
