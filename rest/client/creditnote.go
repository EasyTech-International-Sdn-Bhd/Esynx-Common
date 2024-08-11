package client

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-common/rest/request"
)

const (
	EndPointCreditNote      = "creditnote"
	EndPointCreditNoteBulk  = "creditnote/bulk"
	EndPointCreditNoteQuery = "creditnote/query"
)

type CreditNoteResponse struct {
	Data    []entities.CmsCreditnote `json:"data"`
	Status  bool                     `json:"status"`
	Message string                   `json:"message"`
}

func (i *ApiClient) PostCreditNote(form request.CreditNoteCreateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Post(EndPointCreditNote)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) PostBulkCreditNote(form request.CreditNoteBulkCreateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Post(EndPointCreditNoteBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) GetCreditNote(creditNoteCode string) (*CreditNoteResponse, error) {
	var response CreditNoteResponse
	resp, err := i.Reqwest.R().
		SetQueryParam("cnCode", creditNoteCode).
		SetResult(&response).
		Get(EndPointCreditNote)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) GetBulkCreditNote(creditNoteCodes string) (*CreditNoteResponse, error) {
	var response CreditNoteResponse
	resp, err := i.Reqwest.R().
		SetQueryParam("cnCode", creditNoteCodes).
		SetResult(&response).
		Get(EndPointCreditNoteBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) PutCreditNote(form request.CreditNoteUpdateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Put(EndPointCreditNote)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) PutBulkCreditNote(form request.CreditNoteBulkUpdateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Put(EndPointCreditNoteBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) DeleteCreditNote(form request.CreditNoteDeleteForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Delete(EndPointCreditNote)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) DeleteBulkCreditNote(form request.CreditNoteBulkDeleteForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Delete(EndPointCreditNoteBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) QueryCreditNote(query interface{}) (*CreditNoteResponse, error) {
	var response CreditNoteResponse
	resp, err := i.Reqwest.R().
		SetBody(query).
		SetResult(&response).
		Post(EndPointCreditNoteQuery)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}
