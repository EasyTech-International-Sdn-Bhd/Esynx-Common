package rest

import "github.com/easytech-international-sdn-bhd/esynx-common/entities"

type ReferenceBasedResponse struct {
	Data    []string `json:"data"`
	Status  bool     `json:"status"`
	Message string   `json:"message"`
}

type InvoiceResponse struct {
	Data    []entities.CmsInvoice `json:"data"`
	Status  bool                  `json:"status"`
	Message string                `json:"message"`
}

type InvoiceSalesResponse struct {
	Data    []entities.CmsInvoiceSales `json:"data"`
	Status  bool                       `json:"status"`
	Message string                     `json:"message"`
}

type InvoiceDetailsResponse struct {
	Data    []entities.CmsInvoiceDetails `json:"data"`
	Status  bool                         `json:"status"`
	Message string                       `json:"message"`
}

type DebitNoteResponse struct {
	Data    []entities.CmsDebitnote `json:"data"`
	Status  bool                    `json:"status"`
	Message string                  `json:"message"`
}

type DebitNoteSalesResponse struct {
	Data    []entities.CmsDebitnoteSales `json:"data"`
	Status  bool                         `json:"status"`
	Message string                       `json:"message"`
}

type DebitNoteDetailsResponse struct {
	Data    []entities.CmsDebitnoteDetails `json:"data"`
	Status  bool                           `json:"status"`
	Message string                         `json:"message"`
}

type CreditNoteResponse struct {
	Data    []entities.CmsCreditnote `json:"data"`
	Status  bool                     `json:"status"`
	Message string                   `json:"message"`
}

type CreditNoteSalesResponse struct {
	Data    []entities.CmsCreditnoteSales `json:"data"`
	Status  bool                          `json:"status"`
	Message string                        `json:"message"`
}

type CreditNoteDetailsResponse struct {
	Data    []entities.CmsCreditnoteDetails `json:"data"`
	Status  bool                            `json:"status"`
	Message string                          `json:"message"`
}
