package forms

import (
	"time"
)

type CmsDebitNoteForm struct {
	DnCode            string    `xorm:"index unique VARCHAR(20)" json:"dnCode,omitempty" xml:"dnCode" validate:"trim"`
	CustCode          string    `xorm:"index VARCHAR(20)" json:"custCode,omitempty" xml:"custCode" validate:"trim"`
	DnDate            time.Time `xorm:"TIMESTAMP" json:"dnDate,omitempty" xml:"dnDate"`
	DnAmount          float64   `xorm:"DOUBLE" json:"dnAmount,omitempty" xml:"dnAmount"`
	OutstandingAmount float64   `xorm:"DOUBLE" json:"outstandingAmount,omitempty" xml:"outstandingAmount"`
	Agent             string    `xorm:"VARCHAR(50)" json:"agent,omitempty" xml:"agent" validate:"trim"`
	Cancelled         string    `xorm:"CHAR(1)" json:"cancelled,omitempty" xml:"cancelled" validate:"trim"`
	Approved          int       `xorm:"default 0 INT" json:"approved,omitempty" xml:"approved"`
	Approver          string    `xorm:"VARCHAR(200)" json:"approver,omitempty" xml:"approver" validate:"trim"`
	ApprovedAt        time.Time `xorm:"DATETIME" json:"approvedAt,omitempty" xml:"approvedAt"`
	FromDoc           string    `xorm:"default 'SL' ENUM('AR','SL')" json:"fromDoc,omitempty" xml:"fromDoc" validate:"trim"`
	UpdatedAt         time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
	RefNo             string    `xorm:"VARCHAR(80)" json:"refNo,omitempty" xml:"refNo" validate:"trim"`
	DnUdf             string    `json:"dnUdf,omitempty" xml:"invUdf" validate:"trim"`
}

type CmsDebitNoteDetailsForm struct {
	DnCode       string  `xorm:"VARCHAR(100)" json:"dnCode,omitempty" xml:"dnCode" validate:"trim"`
	ItemCode     string  `xorm:"VARCHAR(200)" json:"itemCode,omitempty" xml:"itemCode" validate:"trim"`
	ItemName     string  `xorm:"VARCHAR(200)" json:"itemName,omitempty" xml:"itemName" validate:"trim"`
	ItemPrice    float64 `xorm:"default 0 DOUBLE" json:"itemPrice,omitempty" xml:"itemPrice"`
	Quantity     float64 `xorm:"default 0 DOUBLE" json:"quantity,omitempty" xml:"quantity"`
	TotalPrice   float64 `xorm:"default 0 DOUBLE" json:"totalPrice,omitempty" xml:"totalPrice"`
	Uom          string  `xorm:"VARCHAR(200)" json:"uom,omitempty" xml:"uom" validate:"trim"`
	Discount     string  `xorm:"comment('0%+10+50%') VARCHAR(100)" json:"discount,omitempty" xml:"discount" validate:"trim"`
	ActiveStatus int     `xorm:"default 1 comment('0=inactive, 1=active') INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	RefNo        string  `xorm:"unique VARCHAR(200)" json:"refNo,omitempty" xml:"refNo" validate:"trim"`
	DnDtlUdf     string  `xorm:"JSON" json:"dnDtlUdf,omitempty" xml:"dnDtlUdf" validate:"trim"`
}
