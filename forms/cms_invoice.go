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
	Agent             string    `json:"agent,omitempty" xml:"agent"`
	InvUdf            string    `json:"invUdf,omitempty" xml:"invUdf"`
	RefNo             string    `json:"refNo,omitempty" xml:"refNo"`
	DocType           string    `json:"docType,omitempty" xml:"docType"`
	FromDoc           string    `json:"fromDoc,omitempty" xml:"fromDoc"`
	Cancelled         string    `json:"cancelled,omitempty" xml:"cancelled"`
	Termcode          string    `json:"termcode,omitempty" xml:"termcode"`
}

type CmsInvoiceDetailsForm struct {
	InvoiceCode  string  `xorm:"not null unique(invoice_code) VARCHAR(50)" json:"invoiceCode,omitempty" xml:"invoiceCode"`
	ItemCode     string  `xorm:"not null VARCHAR(50)" json:"itemCode,omitempty" xml:"itemCode"`
	ItemName     string  `xorm:"not null VARCHAR(200)" json:"itemName,omitempty" xml:"itemName"`
	ItemPrice    float64 `xorm:"default 0 DOUBLE" json:"itemPrice,omitempty" xml:"itemPrice"`
	Quantity     float64 `xorm:"default 0 DOUBLE" json:"quantity,omitempty" xml:"quantity"`
	TotalPrice   float64 `xorm:"default 0 DOUBLE" json:"totalPrice,omitempty" xml:"totalPrice"`
	Uom          string  `xorm:"not null VARCHAR(20)" json:"uom,omitempty" xml:"uom"`
	Discount     string  `xorm:"default '0' VARCHAR(50)" json:"discount,omitempty" xml:"discount"`
	ActiveStatus int     `xorm:"default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	SequenceNo   int     `xorm:"not null default 0 INT" json:"sequenceNo,omitempty" xml:"sequenceNo"`
	InvDtlUdf    string  `xorm:"JSON" json:"invDtlUdf,omitempty" xml:"invDtlUdf"`
	RefNo        string  `xorm:"unique(invoice_code) VARCHAR(200)" json:"refNo,omitempty" xml:"refNo"`
	SessionId    string  `xorm:"default '' VARCHAR(100)" json:"sessionId,omitempty" xml:"sessionId"`
}
