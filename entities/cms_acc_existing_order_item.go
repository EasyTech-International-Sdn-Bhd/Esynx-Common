package entities

import (
	"time"
)

type CmsAccExistingOrderItem struct {
	OrderItemId         uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"orderItemId,omitempty" xml:"orderItemId"`
	OrderId             string    `xorm:"index unique(unx) VARCHAR(20)" json:"orderId,omitempty" xml:"orderId"`
	IpadItemId          int       `xorm:"default 0 unique(unx) INT" json:"ipadItemId,omitempty" xml:"ipadItemId"`
	ProductCode         string    `xorm:"unique(unx) index VARCHAR(50)" json:"productCode,omitempty" xml:"productCode"`
	ProductName         string    `xorm:"VARCHAR(400)" json:"productName,omitempty" xml:"productName"`
	AgentRemark         string    `xorm:"not null default '' VARCHAR(200)" json:"agentRemark,omitempty" xml:"agentRemark"`
	Quantity            float64   `xorm:"DOUBLE" json:"quantity,omitempty" xml:"quantity"`
	EdittedQuantity     string    `xorm:"VARCHAR(50)" json:"edittedQuantity,omitempty" xml:"edittedQuantity"`
	UnitPrice           float64   `xorm:"DOUBLE" json:"unitPrice,omitempty" xml:"unitPrice"`
	Disc1               float64   `xorm:"DOUBLE" json:"disc1,omitempty" xml:"disc1"`
	Disc2               float64   `xorm:"DOUBLE" json:"disc2,omitempty" xml:"disc2"`
	Disc3               float64   `xorm:"DOUBLE" json:"disc3,omitempty" xml:"disc3"`
	UnitUom             string    `xorm:"VARCHAR(100)" json:"unitUom,omitempty" xml:"unitUom"`
	AttributeRemark     string    `xorm:"comment('The chosen attribute name and value will be stored here, e.g. Size=L, Colour=Red') BLOB" json:"attributeRemark,omitempty" xml:"attributeRemark"`
	OptionalRemark      string    `xorm:"comment('The selected optional item will options here, e.g. Sport Rim, Leather Seat.') VARCHAR(100)" json:"optionalRemark,omitempty" xml:"optionalRemark"`
	PickerNote          string    `xorm:"not null VARCHAR(500)" json:"pickerNote,omitempty" xml:"pickerNote"`
	DiscountAmount      string    `xorm:"VARCHAR(30)" json:"discountAmount,omitempty" xml:"discountAmount"`
	SubTotal            float64   `xorm:"comment('the optional item price will affect the sub-total') DOUBLE" json:"subTotal,omitempty" xml:"subTotal"`
	SequenceNo          int       `xorm:"INT" json:"sequenceNo,omitempty" xml:"sequenceNo"`
	PackingStatus       int       `xorm:"default 0 comment('0=not packed, 1=packed, 2=no stock, 3=no stock but informed') INT" json:"packingStatus,omitempty" xml:"packingStatus"`
	PackedQty           float32   `xorm:"default 0 FLOAT" json:"packedQty,omitempty" xml:"packedQty"`
	Isparent            int       `xorm:"INT" json:"isparent,omitempty" xml:"isparent"`
	ParentCode          string    `xorm:"VARCHAR(100)" json:"parentCode,omitempty" xml:"parentCode"`
	PackedBy            string    `xorm:"VARCHAR(200)" json:"packedBy,omitempty" xml:"packedBy"`
	PackConfirmedBy     string    `xorm:"VARCHAR(50)" json:"packConfirmedBy,omitempty" xml:"packConfirmedBy"`
	PackConfirmedStatus int       `xorm:"default 0 INT" json:"packConfirmedStatus,omitempty" xml:"packConfirmedStatus"`
	OrderItemValidity   int       `xorm:"default 2 comment('0 = reject 1=pending 2 = approved') INT" json:"orderItemValidity,omitempty" xml:"orderItemValidity"`
	CancelStatus        int       `xorm:"default 0 comment('0=not canceled, 1=canceld') INT" json:"cancelStatus,omitempty" xml:"cancelStatus"`
	UpdatedAt           time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
	BatchNo             []byte    `xorm:"JSON" json:"batchNo,omitempty" xml:"batchNo"`
	SerialNo            []byte    `xorm:"JSON" json:"serialNo,omitempty" xml:"serialNo"`
	ItemUdf             string    `xorm:"JSON" json:"itemUdf,omitempty" xml:"itemUdf"`
}

func (m *CmsAccExistingOrderItem) TableName() string {
	return "cms_acc_existing_order_item"
}
