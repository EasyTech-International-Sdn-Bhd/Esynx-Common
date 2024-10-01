package entities

import (
	"time"
)

type CmsOrderItem struct {
	OrderItemId         uint64    `xorm:"pk autoincr unique(unique_key) UNSIGNED BIGINT" json:"orderItemId,omitempty" xml:"orderItemId"`
	OrderId             string    `xorm:"unique(unique_key) index VARCHAR(20)" json:"orderId,omitempty" xml:"orderId"`
	IpadItemId          int64     `xorm:"default 0 unique(unique_key) BIGINT" json:"ipadItemId,omitempty" xml:"ipadItemId"`
	ProductCode         string    `xorm:"VARCHAR(50) index" json:"productCode,omitempty" xml:"productCode"`
	ProductName         string    `xorm:"VARCHAR(200)" json:"productName,omitempty" xml:"productName"`
	AgentRemark         string    `xorm:"not null BLOB" json:"agentRemark,omitempty" xml:"agentRemark"`
	Quantity            float64   `xorm:"DOUBLE" json:"quantity,omitempty" xml:"quantity"`
	EdittedQuantity     string    `xorm:"VARCHAR(50)" json:"edittedQuantity,omitempty" xml:"edittedQuantity"`
	UnitPrice           float64   `xorm:"DOUBLE" json:"unitPrice,omitempty" xml:"unitPrice"`
	Disc1               float64   `xorm:"default 0 DOUBLE" json:"disc1,omitempty" xml:"disc1"`
	Disc2               float64   `xorm:"default 0 DOUBLE" json:"disc2,omitempty" xml:"disc2"`
	Disc3               float64   `xorm:"default 0 DOUBLE" json:"disc3,omitempty" xml:"disc3"`
	UnitUom             string    `xorm:"VARCHAR(100)" json:"unitUom,omitempty" xml:"unitUom"`
	AttributeRemark     string    `xorm:"comment('The chosen attribute name and value will be stored here, e.g. Size=L, Colour=Red') BLOB" json:"attributeRemark,omitempty" xml:"attributeRemark"`
	OptionalRemark      string    `xorm:"comment('The selected optional item will options here, e.g. Sport Rim, Leather Seat.') BLOB" json:"optionalRemark,omitempty" xml:"optionalRemark"`
	PickerNote          string    `xorm:"not null VARCHAR(500)" json:"pickerNote,omitempty" xml:"pickerNote"`
	DiscountAmount      string    `xorm:"default '' VARCHAR(50)" json:"discountAmount,omitempty" xml:"discountAmount"`
	SubTotal            float64   `xorm:"comment('the optional item price will affect the sub-total') DOUBLE" json:"subTotal,omitempty" xml:"subTotal"`
	SequenceNo          int       `xorm:"INT" json:"sequenceNo,omitempty" xml:"sequenceNo"`
	PackingStatus       int       `xorm:"default 0 comment('0=not packed, 1=packed, 2=no stock, 3=no stock but informed') INT" json:"packingStatus,omitempty" xml:"packingStatus"`
	PackedQty           int       `xorm:"default 0 INT" json:"packedQty,omitempty" xml:"packedQty"`
	PackConfirmedBy     string    `xorm:"VARCHAR(45)" json:"packConfirmedBy,omitempty" xml:"packConfirmedBy"`
	PackConfirmedStatus int       `xorm:"default 0 INT" json:"packConfirmedStatus,omitempty" xml:"packConfirmedStatus"`
	Isparent            int       `xorm:"INT" json:"isparent,omitempty" xml:"isparent"`
	ParentCode          string    `xorm:"VARCHAR(100)" json:"parentCode,omitempty" xml:"parentCode"`
	PackedBy            string    `xorm:"VARCHAR(200)" json:"packedBy,omitempty" xml:"packedBy"`
	OrderItemValidity   int       `xorm:"default 2 comment('0 = reject 1=pending 2 = approved') INT" json:"orderItemValidity,omitempty" xml:"orderItemValidity"`
	CancelStatus        int       `xorm:"default 0 comment('0=not canceled, 1=canceld') INT" json:"cancelStatus,omitempty" xml:"cancelStatus"`
	UpdatedAt           time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
	WarehouseCode       string    `xorm:"VARCHAR(100)" json:"warehouseCode,omitempty" xml:"warehouseCode"`
	ProjNo              []byte    `xorm:"JSON" json:"projNo,omitempty" xml:"projNo"`
	IsExchange          int       `xorm:"not null default 0 INT" json:"isExchange,omitempty" xml:"isExchange"`
	UnitLength          []byte    `xorm:"JSON" json:"unitLength,omitempty" xml:"unitLength"`
	IsReturn            int       `xorm:"not null default 0 TINYINT(1)" json:"isReturn,omitempty" xml:"isReturn"`
	ReturnItem          []byte    `xorm:"JSON" json:"returnItem,omitempty" xml:"returnItem"`
	LastUpdatedBy       string    `xorm:"VARCHAR(100)" json:"lastUpdatedBy,omitempty" xml:"lastUpdatedBy"`
}

func (m *CmsOrderItem) TableName() string {
	return "cms_order_item"
}
