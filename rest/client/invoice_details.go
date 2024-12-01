package client

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-common/rest/request"
)

const (
	EndPointInvoiceDetails      = "invoice/details"
	EndPointInvoiceDetailsBulk  = "invoice/details/bulk"
	EndPointInvoiceDetailsQuery = "invoice/details/query"
)

type InvoiceDetailsResponse struct {
	Data    []entities.CmsInvoiceDetails `json:"data"`
	Status  bool                         `json:"status"`
	Message string                       `json:"message"`
}

func (i *ApiClient) PostInvoiceDetails(form request.InvoiceDetailsCreateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Post(EndPointInvoiceDetails)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) PostBulkInvoiceDetails(form request.InvoiceDetailsBulkCreateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Post(EndPointInvoiceDetailsBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) GetInvoiceDetails(invoiceCode string) (*InvoiceDetailsResponse, error) {
	var response InvoiceDetailsResponse
	resp, err := i.Reqwest.R().
		SetPathParam("invoiceCode", invoiceCode).
		SetResult(&response).
		Get(fmt.Sprintf("%s/{invoiceCode}", EndPointInvoiceDetails))

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) GetBulkInvoiceDetails(invoiceCodes string) (*InvoiceDetailsResponse, error) {
	var response InvoiceDetailsResponse
	resp, err := i.Reqwest.R().
		SetPathParam("invoiceCode", invoiceCodes).
		SetResult(&response).
		Get(fmt.Sprintf("%s/{invoiceCode}", EndPointInvoiceDetailsBulk))

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) PutInvoiceDetails(form request.InvoiceDetailsUpdateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Put(EndPointInvoiceDetails)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) PutBulkInvoiceDetails(form request.InvoiceDetailsBulkUpdateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Put(EndPointInvoiceDetailsBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) DeleteInvoiceDetails(form request.InvoiceDetailsDeleteForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Delete(EndPointInvoiceDetails)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) DeleteBulkInvoiceDetails(form request.InvoiceDetailsBulkDeleteForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Delete(EndPointInvoiceDetailsBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) DeleteInvoiceDetailsByQuery(query interface{}) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(query).
		SetResult(&response).
		Delete(EndPointInvoiceDetailsQuery)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}
	return &response, nil
}

func (i *ApiClient) QueryInvoiceDetails(query interface{}) (*InvoiceDetailsResponse, error) {
	var response InvoiceDetailsResponse
	resp, err := i.Reqwest.R().
		SetBody(query).
		SetResult(&response).
		Post(EndPointInvoiceDetailsQuery)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}
