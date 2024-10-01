package entities

import "time"

type CmsDebitnoteDetails struct {
	Id           uint64    `xorm:"pk autoincr unique(dn_code) UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	DnCode       string    `xorm:"not null unique(dn_code) index VARCHAR(40)" json:"dnCode,omitempty" xml:"dnCode"`
	ItemCode     string    `xorm:"not null VARCHAR(100)" json:"itemCode,omitempty" xml:"itemCode"`
	ItemName     string    `xorm:"not null VARCHAR(200)" json:"itemName,omitempty" xml:"itemName"`
	ItemPrice    float64   `xorm:"default 0 DOUBLE" json:"itemPrice,omitempty" xml:"itemPrice"`
	Quantity     float64   `xorm:"default 0 DOUBLE" json:"quantity,omitempty" xml:"quantity"`
	Uom          string    `xorm:"not null VARCHAR(200)" json:"uom,omitempty" xml:"uom"`
	TotalPrice   float64   `xorm:"default 0 DOUBLE" json:"totalPrice,omitempty" xml:"totalPrice"`
	Discount     string    `xorm:"default '0' comment('0%+10+50%') VARCHAR(100)" json:"discount,omitempty" xml:"discount"`
	SequenceNo   int       `xorm:"not null default 0 INT" json:"sequenceNo,omitempty" xml:"sequenceNo"`
	ActiveStatus int       `xorm:"default 1 INT comment('0=inactive, 1=active')" json:"activeStatus,omitempty" xml:"activeStatus"`
	UpdatedAt    time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
	RefNo        string    `xorm:"unique(dn_code) index VARCHAR(200)" json:"refNo,omitempty" xml:"refNo"`
	DnDtlUdf     string    `xorm:"JSON" json:"dnDtlUdf,omitempty" xml:"dnDtlUdf"`
}

func (m *CmsDebitnoteDetails) TableName() string {
	return "cms_debitnote_details"
}

func (m *CmsDebitnoteDetails) BeforeInsert() {
	m.BeforeUpdate()
}

func (m *CmsDebitnoteDetails) BeforeUpdate() {
	m.UpdatedAt = time.Now()
}
