package entities

import (
	"time"
)

type CmsProductPrice struct {
	ProductPriceId uint64    `xorm:"pk autoincr UNSIGNED BIGINT" json:"productPriceId,omitempty" xml:"productPriceId"`
	ProductCode    string    `xorm:"unique(xunique) VARCHAR(100) index" json:"productCode,omitempty" xml:"productCode"`
	PriceCat       string    `xorm:"unique(xunique) VARCHAR(50)" json:"priceCat,omitempty" xml:"priceCat"`
	CustCode       string    `xorm:"VARCHAR(30) index" json:"custCode,omitempty" xml:"custCode"`
	ProductPrice   float64   `xorm:"default 0 DOUBLE" json:"productPrice,omitempty" xml:"productPrice"`
	Disc1          float64   `xorm:"default 0 DOUBLE" json:"disc1,omitempty" xml:"disc1"`
	Disc2          float64   `xorm:"default 0 DOUBLE" json:"disc2,omitempty" xml:"disc2"`
	Disc3          float64   `xorm:"default 0 DOUBLE" json:"disc3,omitempty" xml:"disc3"`
	DateFrom       time.Time `xorm:"DATETIME" json:"dateFrom,omitempty" xml:"dateFrom"`
	DateTo         time.Time `xorm:"DATETIME" json:"dateTo,omitempty" xml:"dateTo"`
	ProductUom     string    `xorm:"unique(unique) VARCHAR(20)" json:"productUom,omitempty" xml:"productUom"`
	PriceRemark    string    `xorm:"VARCHAR(50)" json:"priceRemark,omitempty" xml:"priceRemark"`
	ActiveStatus   int       `xorm:"default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	Quantity       float64   `xorm:"default 1 DOUBLE" json:"quantity,omitempty" xml:"quantity"`
	UpdatedAt      time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
	PriceUdf       string    `xorm:"JSON" json:"PriceUdf,omitempty" xml:"PriceUdf"`
	RefNo          string    `xorm:"VARCHAR(100)" json:"refNo,omitempty" xml:"refNo"`
	MixDisc        string    `xorm:"VARCHAR(100)" json:"mixDisc,omitempty" xml:"mixDisc"`
}

func (m *CmsProductPrice) TableName() string {
	return "cms_product_price"
}

func (m *CmsProductPrice) BeforeInsert() {
	m.BeforeUpdate()
}

func (m *CmsProductPrice) BeforeUpdate() {
	m.UpdatedAt = time.Now()
	if m.DateFrom.IsZero() {
		m.DateFrom = time.Now().AddDate(-1, 0, 0)
	}
	if m.DateTo.IsZero() {
		m.DateTo = time.Now().AddDate(10, 0, 0)
	}
}
