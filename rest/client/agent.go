package client

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-common/rest/request"
)

const (
	EndPointAgent      = "agent"
	EndPointAgentBulk  = "agent/bulk"
	EndPointAgentQuery = "agent/query"
)

type AgentListResponse struct {
	Data    []entities.CmsLogin `json:"data"`
	Status  bool                `json:"status"`
	Message string              `json:"message"`
}

type AgentResponse struct {
	Data    entities.CmsLogin `json:"data"`
	Status  bool              `json:"status"`
	Message string            `json:"message"`
}

func (i *ApiClient) PostAgent(form request.AgentCreateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Post(EndPointAgent)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) PostBulkAgent(form request.AgentBulkCreateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Post(EndPointAgentBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) GetAgent(agentCode string) (*AgentResponse, error) {
	var response AgentResponse
	resp, err := i.Reqwest.R().
		SetQueryParam("agentCode", agentCode).
		SetResult(&response).
		Get(EndPointAgent)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) GetBulkAgent(agentCodes string) (*AgentListResponse, error) {
	var response AgentListResponse
	resp, err := i.Reqwest.R().
		SetQueryParam("agentCode", agentCodes).
		SetResult(&response).
		Get(EndPointAgentBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) PutAgent(form request.AgentCreateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Put(EndPointAgent)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) PutBulkAgent(form request.AgentBulkUpdateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Put(EndPointAgentBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) DeleteAgent(form request.AgentDeleteForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Delete(EndPointAgent)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) DeleteBulkAgent(form request.AgentBulkDeleteForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Delete(EndPointAgentBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) QueryAgent(query interface{}) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(query).
		SetResult(&response).
		Post(EndPointAgentQuery)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}
