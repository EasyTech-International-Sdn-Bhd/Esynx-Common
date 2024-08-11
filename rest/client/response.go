package client

import "github.com/easytech-international-sdn-bhd/esynx-common/entities"

type ReferenceBasedResponse struct {
	Data    []string `json:"data"`
	Status  bool     `json:"status"`
	Message string   `json:"message"`
}

type InvoiceSalesResponse struct {
	Data    []entities.CmsInvoiceSales `json:"data"`
	Status  bool                       `json:"status"`
	Message string                     `json:"message"`
}

type DebitNoteSalesResponse struct {
	Data    []entities.CmsDebitnoteSales `json:"data"`
	Status  bool                         `json:"status"`
	Message string                       `json:"message"`
}

type CreditNoteSalesResponse struct {
	Data    []entities.CmsCreditnoteSales `json:"data"`
	Status  bool                          `json:"status"`
	Message string                        `json:"message"`
}
