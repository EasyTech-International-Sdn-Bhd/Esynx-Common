package entities

import (
	"time"
)

type CmsProductUomPrice struct {
	ProductUomPriceId   uint64    `xorm:"pk autoincr UNSIGNED BIGINT" json:"productUomPriceId,omitempty" xml:"productUomPriceId"`
	ProductCode         string    `xorm:"unique(unique_uom) VARCHAR(100) index" json:"productCode,omitempty" xml:"productCode"`
	ProductUom          string    `xorm:"unique(unique_uom) VARCHAR(20)" json:"productUom,omitempty" xml:"productUom"`
	ProductUomRate      float64   `xorm:"default 0 unique(unique_uom) DOUBLE" json:"productUomRate,omitempty" xml:"productUomRate"`
	ProductStdPrice     float64   `xorm:"default 0 DOUBLE" json:"productStdPrice,omitempty" xml:"productStdPrice"`
	ProductMinPrice     float64   `xorm:"default 0 DOUBLE" json:"productMinPrice,omitempty" xml:"productMinPrice"`
	ProductDefaultPrice int       `xorm:"default 0 comment('Each product only can select 1 default price, 0=not default, 1=default') INT" json:"productDefaultPrice,omitempty" xml:"productDefaultPrice"`
	ActiveStatus        int       `xorm:"default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	UpdatedAt           time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
	PriceUdf            string    `xorm:"JSON" json:"priceUdf,omitempty" xml:"priceUdf"`
	QrCode              string    `xorm:"VARCHAR(30)" json:"qrCode,omitempty" xml:"qrCode"`
	RefNo               string    `xorm:"VARCHAR(200)" json:"refNo,omitempty" xml:"refNo"`
}

func (m *CmsProductUomPrice) TableName() string {
	return "cms_product_uom_price"
}

func (m *CmsProductUomPrice) BeforeInsert() {
	m.BeforeUpdate()
}

func (m *CmsProductUomPrice) BeforeUpdate() {
	m.UpdatedAt = time.Now()
}
