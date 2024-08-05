package forms

import "time"

type CmsInvoiceForm struct {
	InvoiceCode       string    `json:"invoiceCode,omitempty" xml:"invoiceCode"`
	CustCode          string    `json:"custCode,omitempty" xml:"custCode"`
	InvoiceDate       time.Time `json:"invoiceDate,omitempty" xml:"invoiceDate"`
	InvoiceDueDate    time.Time `json:"invoiceDueDate,omitempty" xml:"invoiceDueDate"`
	InvoiceAmount     float64   `json:"invoiceAmount,omitempty" xml:"invoiceAmount"`
	OutstandingAmount float64   `json:"outstandingAmount,omitempty" xml:"outstandingAmount"`
	Approved          int       `json:"approved,omitempty" xml:"approved"`
	Approver          string    `json:"approver,omitempty" xml:"approver"`
	ApprovedAt        time.Time `json:"approvedAt,omitempty" xml:"approvedAt"`
	Agent             string    `json:"salespersonId,omitempty" xml:"salespersonId"`
	InvUdf            string    `json:"invUdf,omitempty" xml:"invUdf"`
	RefNo             string    `json:"refNo,omitempty" xml:"refNo"`
	DocType           string    `json:"docType,omitempty" xml:"docType"`
	FromDoc           string    `json:"fromDoc,omitempty" xml:"fromDoc"`
	Cancelled         string    `json:"cancelled,omitempty" xml:"cancelled"`
	Termcode          string    `json:"termcode,omitempty" xml:"termcode"`
}
