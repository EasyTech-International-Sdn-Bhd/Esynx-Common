package client

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-common/rest/request"
)

const (
	EndPointDebitNoteDetails      = "debitnote/details"
	EndPointDebitNoteDetailsBulk  = "debitnote/details/bulk"
	EndPointDebitNoteDetailsQuery = "debitnote/details/query"
)

type DebitNoteDetailsResponse struct {
	Data    []entities.CmsDebitnoteDetails `json:"data"`
	Status  bool                           `json:"status"`
	Message string                         `json:"message"`
}

func (i *ApiClient) PostDebitNoteDetails(form request.DebitNoteDetailsCreateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Post(EndPointDebitNoteDetails)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) PostBulkDebitNoteDetails(form request.DebitNoteDetailsBulkCreateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Post(EndPointDebitNoteDetailsBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) GetDebitNoteDetails(dnCode string) (*DebitNoteDetailsResponse, error) {
	var response DebitNoteDetailsResponse
	resp, err := i.Reqwest.R().
		SetQueryParam("dnCode", dnCode).
		SetResult(&response).
		Get(EndPointDebitNoteDetails)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) GetBulkDebitNoteDetails(dnCodes string) (*DebitNoteDetailsResponse, error) {
	var response DebitNoteDetailsResponse
	resp, err := i.Reqwest.R().
		SetQueryParam("dnCode", dnCodes).
		SetResult(&response).
		Get(EndPointDebitNoteDetailsBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) PutDebitNoteDetails(form request.DebitNoteDetailsUpdateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Put(EndPointDebitNoteDetails)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) PutBulkDebitNoteDetails(form request.DebitNoteDetailsBulkUpdateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Put(EndPointDebitNoteDetailsBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) DeleteDebitNoteDetails(form request.DebitNoteDetailsDeleteForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Delete(EndPointDebitNoteDetails)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) DeleteBulkDebitNoteDetails(form request.DebitNoteDetailsBulkDeleteForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Delete(EndPointDebitNoteDetailsBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) QueryDebitNoteDetails(query interface{}) (*DebitNoteDetailsResponse, error) {
	var response DebitNoteDetailsResponse
	resp, err := i.Reqwest.R().
		SetBody(query).
		SetResult(&response).
		Post(EndPointDebitNoteDetailsQuery)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}
