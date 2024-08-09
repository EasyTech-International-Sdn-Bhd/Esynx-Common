package forms

import "time"

type CmsDebitNoteForm struct {
	DnCode            string    `xorm:"index unique VARCHAR(20)" json:"dnCode,omitempty" xml:"dnCode"`
	CustCode          string    `xorm:"index VARCHAR(20)" json:"custCode,omitempty" xml:"custCode"`
	DnDate            time.Time `xorm:"TIMESTAMP" json:"dnDate,omitempty" xml:"dnDate"`
	DnAmount          float64   `xorm:"DOUBLE" json:"dnAmount,omitempty" xml:"dnAmount"`
	OutstandingAmount float64   `xorm:"DOUBLE" json:"outstandingAmount,omitempty" xml:"outstandingAmount"`
	Agent             string    `xorm:"VARCHAR(50)" json:"agent,omitempty" xml:"agent"`
	Cancelled         string    `xorm:"CHAR(1)" json:"cancelled,omitempty" xml:"cancelled"`
	Approved          int       `xorm:"default 0 INT" json:"approved,omitempty" xml:"approved"`
	Approver          string    `xorm:"VARCHAR(200)" json:"approver,omitempty" xml:"approver"`
	ApprovedAt        time.Time `xorm:"DATETIME" json:"approvedAt,omitempty" xml:"approvedAt"`
	FromDoc           string    `xorm:"default 'SL' ENUM('AR','SL')" json:"fromDoc,omitempty" xml:"fromDoc"`
	UpdatedAt         time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
	RefNo             string    `xorm:"VARCHAR(80)" json:"refNo,omitempty" xml:"refNo"`
	DnUdf             string    `json:"dnUdf,omitempty" xml:"invUdf"`
}

type CmsDebitNoteDetailsForm struct {
	DnCode       string `xorm:"VARCHAR(100)" json:"dnCode,omitempty" xml:"dnCode"`
	ItemCode     string `xorm:"VARCHAR(200)" json:"itemCode,omitempty" xml:"itemCode"`
	ItemName     string `xorm:"VARCHAR(200)" json:"itemName,omitempty" xml:"itemName"`
	ItemPrice    string `xorm:"VARCHAR(200)" json:"itemPrice,omitempty" xml:"itemPrice"`
	Quantity     string `xorm:"VARCHAR(200)" json:"quantity,omitempty" xml:"quantity"`
	Uom          string `xorm:"VARCHAR(200)" json:"uom,omitempty" xml:"uom"`
	TotalPrice   string `xorm:"VARCHAR(200)" json:"totalPrice,omitempty" xml:"totalPrice"`
	Discount     string `xorm:"comment('0%+10+50%') VARCHAR(100)" json:"discount,omitempty" xml:"discount"`
	ActiveStatus int    `xorm:"default 1 comment('0=inactive, 1=active') INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	RefNo        string `xorm:"unique VARCHAR(200)" json:"refNo,omitempty" xml:"refNo"`
	DnDtlUdf     string `xorm:"JSON" json:"dnDtlUdf,omitempty" xml:"dnDtlUdf"`
}
