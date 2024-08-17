package forms

import (
	"time"
)

type CmsCreditNoteForm struct {
	CnCode           string    `xorm:"index unique VARCHAR(20)" json:"cnCode,omitempty" xml:"cnCode" validate:"trim"`
	CustCode         string    `xorm:"index VARCHAR(20)" json:"custCode,omitempty" xml:"custCode" validate:"trim"`
	CnDate           time.Time `xorm:"TIMESTAMP" json:"cnDate,omitempty" xml:"cnDate"`
	CnUdf            string    `xorm:"not null JSON" json:"cnUdf,omitempty" xml:"cnUdf"`
	CnAmount         float64   `xorm:"DOUBLE" json:"cnAmount,omitempty" xml:"cnAmount"`
	Agent            string    `json:"agent,omitempty" xml:"agent" validate:"trim"`
	Cancelled        string    `xorm:"CHAR(1)" json:"cancelled,omitempty" xml:"cancelled" validate:"trim"`
	UpdatedAt        time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
	CnKnockoffAmount float64   `xorm:"DOUBLE" json:"cnKnockoffAmount,omitempty" xml:"cnKnockoffAmount"`
	Approved         int       `xorm:"default 0 INT" json:"approved,omitempty" xml:"approved"`
	Approver         string    `xorm:"VARCHAR(200)" json:"approver,omitempty" xml:"approver" validate:"trim"`
	ApprovedAt       time.Time `xorm:"DATETIME" json:"approvedAt,omitempty" xml:"approvedAt"`
	FromDoc          string    `xorm:"default 'SL' ENUM('AR','SL')" json:"fromDoc,omitempty" xml:"fromDoc" validate:"trim"`
	RefNo            string    `xorm:"VARCHAR(80)" json:"refNo,omitempty" xml:"refNo" validate:"trim"`
}

type CmsCreditNoteDetailsForm struct {
	CnCode       string  `xorm:"unique(cn_code) VARCHAR(100)" json:"cnCode,omitempty" xml:"cnCode" validate:"trim"`
	ItemCode     string  `xorm:"unique(cn_code) VARCHAR(200)" json:"itemCode,omitempty" xml:"itemCode" validate:"trim"`
	ItemName     string  `xorm:"VARCHAR(200)" json:"itemName,omitempty" xml:"itemName" validate:"trim"`
	ItemPrice    float64 `xorm:"VARCHAR(200)" json:"itemPrice,omitempty" xml:"itemPrice" validate:"trim"`
	Quantity     float64 `xorm:"default 0 DOUBLE" json:"quantity,omitempty" xml:"quantity"`
	Uom          string  `xorm:"VARCHAR(200)" json:"uom,omitempty" xml:"uom" validate:"trim"`
	TotalPrice   float64 `xorm:"default 0 DOUBLE" json:"totalPrice,omitempty" xml:"totalPrice"`
	Discount     string  `xorm:"comment('0%+10+50%') VARCHAR(100)" json:"discount,omitempty" xml:"discount" validate:"trim"`
	ActiveStatus int     `xorm:"default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	SequenceNo   int     `xorm:"not null default 0 INT" json:"sequenceNo,omitempty" xml:"sequenceNo"`
	CnDtlUdf     string  `xorm:"not null JSON" json:"cnDtlUdf,omitempty" xml:"cnDtlUdf"`
	RefNo        string  `xorm:"unique(cn_code) VARCHAR(200)" json:"refNo,omitempty" xml:"refNo" validate:"trim"`
}
