package forms

import "time"

type CmsInvoiceForm struct {
	InvoiceCode       string    `json:"invoiceCode,omitempty" xml:"invoiceCode" validate:"trim"`
	CustCode          string    `json:"custCode,omitempty" xml:"custCode" validate:"trim"`
	InvoiceDate       time.Time `json:"invoiceDate,omitempty" xml:"invoiceDate"`
	InvoiceDueDate    time.Time `json:"invoiceDueDate,omitempty" xml:"invoiceDueDate"`
	InvoiceAmount     float64   `json:"invoiceAmount,omitempty" xml:"invoiceAmount"`
	OutstandingAmount float64   `json:"outstandingAmount,omitempty" xml:"outstandingAmount"`
	Approved          int       `json:"approved,omitempty" xml:"approved"`
	Approver          string    `json:"approver,omitempty" xml:"approver" validate:"trim"`
	ApprovedAt        time.Time `json:"approvedAt,omitempty" xml:"approvedAt"`
	Agent             string    `json:"agent,omitempty" xml:"agent" validate:"trim"`
	InvUdf            string    `json:"invUdf,omitempty" xml:"invUdf" validate:"trim"`
	RefNo             string    `json:"refNo,omitempty" xml:"refNo" validate:"trim"`
	DocType           string    `json:"docType,omitempty" xml:"docType" validate:"trim"`
	FromDoc           string    `json:"fromDoc,omitempty" xml:"fromDoc" validate:"trim"`
	Cancelled         string    `json:"cancelled,omitempty" xml:"cancelled" validate:"trim"`
	Termcode          string    `json:"termcode,omitempty" xml:"termcode" validate:"trim"`
}

type CmsInvoiceDetailsForm struct {
	InvoiceCode  string  `xorm:"not null unique(invoice_code) VARCHAR(50)" json:"invoiceCode,omitempty" xml:"invoiceCode" validate:"trim"`
	ItemCode     string  `xorm:"not null VARCHAR(80)" json:"itemCode,omitempty" xml:"itemCode" validate:"trim"`
	ItemName     string  `xorm:"not null VARCHAR(250)" json:"itemName,omitempty" xml:"itemName" validate:"trim"`
	ItemPrice    float64 `xorm:"default 0 DOUBLE" json:"itemPrice,omitempty" xml:"itemPrice"`
	Quantity     float64 `xorm:"default 0 DOUBLE" json:"quantity,omitempty" xml:"quantity"`
	TotalPrice   float64 `xorm:"default 0 DOUBLE" json:"totalPrice,omitempty" xml:"totalPrice"`
	Uom          string  `xorm:"not null VARCHAR(20)" json:"uom,omitempty" xml:"uom" validate:"trim"`
	Discount     string  `xorm:"default '0' VARCHAR(50)" json:"discount,omitempty" xml:"discount" validate:"trim"`
	ActiveStatus int     `xorm:"default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	SequenceNo   int     `xorm:"not null default 0 INT" json:"sequenceNo,omitempty" xml:"sequenceNo"`
	InvDtlUdf    string  `xorm:"JSON" json:"invDtlUdf,omitempty" xml:"invDtlUdf" validate:"trim"`
	RefNo        string  `xorm:"unique(invoice_code) VARCHAR(200)" json:"refNo,omitempty" xml:"refNo" validate:"trim"`
}
