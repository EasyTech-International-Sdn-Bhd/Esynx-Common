package request

import "github.com/easytech-international-sdn-bhd/esynx-common/forms"

type InvoiceFormEDA struct {
	Head EdaHeader              `json:"head" binding:"required"`
	Data []forms.CmsInvoiceForm `json:"data" binding:"required"`
}

type InvoiceDetailsFormEDA struct {
	Head EdaHeader                     `json:"head" binding:"required"`
	Data []forms.CmsInvoiceDetailsForm `json:"data" binding:"required"`
}