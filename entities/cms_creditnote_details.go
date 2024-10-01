package entities

import "time"

type CmsCreditnoteDetails struct {
	Id           uint64    `xorm:"pk autoincr unique(cn_code) UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	CnCode       string    `xorm:"not null unique(cn_code) VARCHAR(40)" json:"cnCode,omitempty" xml:"cnCode"`
	ItemCode     string    `xorm:"not null VARCHAR(200)" json:"itemCode,omitempty" xml:"itemCode"`
	ItemName     string    `xorm:"not null VARCHAR(200)" json:"itemName,omitempty" xml:"itemName"`
	ItemPrice    float64   `xorm:"default 0 DOUBLE" json:"itemPrice,omitempty" xml:"itemPrice"`
	Quantity     float64   `xorm:"default 0 DOUBLE" json:"quantity,omitempty" xml:"quantity"`
	Uom          string    `xorm:"not null VARCHAR(200)" json:"uom,omitempty" xml:"uom"`
	TotalPrice   float64   `xorm:"default 0 DOUBLE" json:"totalPrice,omitempty" xml:"totalPrice"`
	Discount     string    `xorm:"default '0' VARCHAR(100)" json:"discount,omitempty" xml:"discount"`
	ActiveStatus int       `xorm:"default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	SequenceNo   int       `xorm:"not null default 0 INT" json:"sequenceNo,omitempty" xml:"sequenceNo"`
	CnDtlUdf     string    `xorm:"not null JSON" json:"cnDtlUdf,omitempty" xml:"cnDtlUdf"`
	RefNo        string    `xorm:"unique(cn_code) index VARCHAR(200)" json:"refNo,omitempty" xml:"refNo"`
	UpdatedAt    time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsCreditnoteDetails) TableName() string {
	return "cms_creditnote_details"
}

func (m *CmsCreditnoteDetails) BeforeInsert() {
	m.BeforeUpdate()
}

func (m *CmsCreditnoteDetails) BeforeUpdate() {
	m.UpdatedAt = time.Now()
}
