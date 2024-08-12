package client

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-common/rest/request"
)

const (
	EndPointDebitNote      = "debitnote"
	EndPointDebitNoteBulk  = "debitnote/bulk"
	EndPointDebitNoteQuery = "debitnote/query"
)

type DebitNoteResponse struct {
	Data    []entities.CmsDebitnote `json:"data"`
	Status  bool                    `json:"status"`
	Message string                  `json:"message"`
}

func (i *ApiClient) PostDebitNote(form request.DebitNoteCreateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Post(EndPointDebitNote)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) PostBulkDebitNote(form request.DebitNoteBulkCreateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Post(EndPointDebitNoteBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) GetDebitNote(dnCode string) (*DebitNoteResponse, error) {
	var response DebitNoteResponse
	resp, err := i.Reqwest.R().
		SetPathParam("dnCode", dnCode).
		SetResult(&response).
		Get(EndPointDebitNote)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) GetBulkDebitNote(dnCodes string) (*DebitNoteResponse, error) {
	var response DebitNoteResponse
	resp, err := i.Reqwest.R().
		SetPathParam("dnCode", dnCodes).
		SetResult(&response).
		Get(EndPointDebitNoteBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) PutDebitNote(form request.DebitNoteUpdateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Put(EndPointDebitNote)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) PutBulkDebitNote(form request.DebitNoteBulkUpdateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Put(EndPointDebitNoteBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) DeleteDebitNote(form request.DebitNoteDeleteForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Delete(EndPointDebitNote)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) DeleteBulkDebitNote(form request.DebitNoteBulkDeleteForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Delete(EndPointDebitNoteBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) QueryDebitNote(query interface{}) (*DebitNoteResponse, error) {
	var response DebitNoteResponse
	resp, err := i.Reqwest.R().
		SetBody(query).
		SetResult(&response).
		Post(EndPointDebitNoteQuery)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}
