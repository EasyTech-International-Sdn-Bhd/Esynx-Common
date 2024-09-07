package request

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/forms"
)

type InvoiceCreateForm struct {
	Reference string               `json:"reference" binding:"required,min=3,max=100"`
	DocType   string               `json:"docType" binding:"required,min=2"`
	Data      forms.CmsInvoiceForm `json:"data" binding:"required"`
}

type InvoiceBulkCreateForm struct {
	Data []InvoiceCreateForm `json:"data" binding:"required,v-nonempty"`
}

type InvoiceUpdateForm struct {
	Reference string               `json:"reference" binding:"required,min=3,max=100"`
	Data      forms.CmsInvoiceForm `json:"data" binding:"required"`
}

type InvoiceBulkUpdateForm struct {
	Data []InvoiceUpdateForm `json:"data" binding:"required,v-nonempty"`
}

type InvoiceDeleteForm struct {
	InvoiceCode string `json:"invoiceCode" binding:"required,min=3,max=100"`
}

type InvoiceBulkDeleteForm struct {
	Data []InvoiceDeleteForm `json:"invoiceCodes" binding:"required,v-nonempty"`
}

type InvoiceDetailsCreateForm struct {
	Reference string                      `json:"reference" binding:"required,min=3,max=100"`
	Data      forms.CmsInvoiceDetailsForm `json:"data" binding:"required"`
}

type InvoiceDetailsBulkCreateForm struct {
	Data []InvoiceDetailsCreateForm `json:"data" binding:"required,v-nonempty"`
}

type InvoiceDetailsUpdateForm struct {
	Reference string                      `json:"reference" binding:"required,min=3,max=100"`
	Data      forms.CmsInvoiceDetailsForm `json:"data" binding:"required"`
}

type InvoiceDetailsBulkUpdateForm struct {
	Data []InvoiceDetailsUpdateForm `json:"data" binding:"required,v-nonempty"`
}

type InvoiceDetailsDeleteForm struct {
	Reference string `json:"id" binding:"required,min=3,max=100"`
}

type InvoiceDetailsBulkDeleteForm struct {
	Data []InvoiceDetailsDeleteForm `json:"ids" binding:"required,v-nonempty"`
}
