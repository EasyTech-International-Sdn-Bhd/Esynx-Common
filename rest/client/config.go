package client

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-common/rest/request"
)

const (
	EndPointEsynxConfig      = "esynx/config"
	EndPointEsynxBulkConfig  = "esynx/config/bulk"
	EndPointEsynxConfigQuery = "esynx/config/query"
)

type EsynxConfigResponse struct {
	Data    []entities.EsynxConfig `json:"data"`
	Status  bool                   `json:"status"`
	Message string                 `json:"message"`
}

func (i *ApiClient) GetBulkEsynxConfig() (*EsynxConfigResponse, error) {
	var response EsynxConfigResponse
	resp, err := i.Reqwest.R().
		SetResult(&response).
		Post(EndPointEsynxBulkConfig)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) PostEsynxConfig(form request.ConfigForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Post(EndPointEsynxConfig)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) PutEsynxConfig(form request.ConfigForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Put(EndPointEsynxConfig)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) DeleteEsynxConfig(form request.ConfigDeleteForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Delete(EndPointEsynxConfig)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) QueryEsynxConfig(query interface{}) (*EsynxConfigResponse, error) {
	var response EsynxConfigResponse
	resp, err := i.Reqwest.R().
		SetBody(query).
		SetResult(&response).
		Post(EndPointEsynxConfigQuery)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}
