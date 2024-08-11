package client

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-common/rest/request"
)

const (
	EndPointInvoice      = "invoice"
	EndPointInvoiceBulk  = "invoice/bulk"
	EndPointInvoiceQuery = "invoice/query"
)

type InvoiceResponse struct {
	Data    []entities.CmsInvoice `json:"data"`
	Status  bool                  `json:"status"`
	Message string                `json:"message"`
}

func (i *ApiClient) PostInvoice(form request.InvoiceCreateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Post(EndPointInvoice)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) PostBulkInvoice(form request.InvoiceBulkCreateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Post(EndPointInvoiceBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) GetInvoice(invoiceCode string) (*InvoiceResponse, error) {
	var response InvoiceResponse
	resp, err := i.Reqwest.R().
		SetQueryParam("invoiceCode", invoiceCode).
		SetResult(&response).
		Get(EndPointInvoice)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) GetBulkInvoice(invoiceCodes string) (*InvoiceResponse, error) {
	var response InvoiceResponse
	resp, err := i.Reqwest.R().
		SetQueryParam("invoiceCode", invoiceCodes).
		SetResult(&response).
		Get(EndPointInvoiceBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) PutInvoice(form request.InvoiceUpdateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Put(EndPointInvoice)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) PutBulkInvoice(form request.InvoiceBulkUpdateForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Put(EndPointInvoiceBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) DeleteInvoice(form request.InvoiceDeleteForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Delete(EndPointInvoice)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) DeleteBulkInvoice(form request.InvoiceBulkDeleteForm) (*ReferenceBasedResponse, error) {
	var response ReferenceBasedResponse
	resp, err := i.Reqwest.R().
		SetBody(form).
		SetResult(&response).
		Delete(EndPointInvoiceBulk)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}

func (i *ApiClient) QueryInvoice(query interface{}) (*InvoiceResponse, error) {
	var response InvoiceResponse
	resp, err := i.Reqwest.R().
		SetBody(query).
		SetResult(&response).
		Post(EndPointInvoiceQuery)

	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}
	return &response, nil
}
