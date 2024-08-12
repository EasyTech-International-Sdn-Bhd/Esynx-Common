package client

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-common/rest/request"
)

const (
	EndPointCreditNoteDetails      = "creditnote/details"
	EndPointCreditNoteDetailsBulk  = "creditnote/details/bulk"
	EndPointCreditNoteDetailsQuery = "creditnote/details/query"
)

type CreditNoteDetailsResponse struct {
	Data    []entities.CmsCreditnoteDetails `json:"data"`
	Status  bool                            `json:"status"`
	Message string                          `json:"message"`
}

func (i *ApiClient) PostCreditNoteDetails(form request.CreditNoteDetailsCreateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Post(EndPointCreditNoteDetails)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) PostBulkCreditNoteDetails(form request.CreditNoteDetailsBulkCreateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Post(EndPointCreditNoteDetailsBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) GetCreditNoteDetails(creditNoteCode string) (*CreditNoteDetailsResponse, error) {
	var response CreditNoteDetailsResponse
	resp, err := i.Reqwest.R().
		SetPathParam("cnCode", creditNoteCode).
		SetResult(&response).
		Get(fmt.Sprintf("%s/{cnCode}", EndPointCreditNoteDetails))

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) GetBulkCreditNoteDetails(creditNoteCodes string) (*CreditNoteDetailsResponse, error) {
	var response CreditNoteDetailsResponse
	resp, err := i.Reqwest.R().
		SetPathParam("cnCode", creditNoteCodes).
		SetResult(&response).
		Get(fmt.Sprintf("%s/{cnCode}", EndPointCreditNoteDetailsBulk))

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) PutCreditNoteDetails(form request.CreditNoteDetailsUpdateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Put(EndPointCreditNoteDetails)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) PutBulkCreditNoteDetails(form request.CreditNoteDetailsBulkUpdateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Put(EndPointCreditNoteDetailsBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) DeleteCreditNoteDetails(form request.CreditNoteDetailsDeleteForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Delete(EndPointCreditNoteDetails)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) DeleteBulkCreditNoteDetails(form request.CreditNoteDetailsBulkDeleteForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Delete(EndPointCreditNoteDetailsBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) QueryCreditNoteDetails(query interface{}) (*CreditNoteDetailsResponse, error) {
	var response CreditNoteDetailsResponse
	resp, err := i.Reqwest.R().
		SetBody(query).
		SetResult(&response).
		Post(EndPointCreditNoteDetailsQuery)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}
