package entities

type CmsStockTransferDtl struct {
	Id           uint64  `xorm:"pk autoincr UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	StCode       string  `xorm:"not null unique(dtl) VARCHAR(40)" json:"stCode,omitempty" xml:"stCode"`
	ProductCode  string  `xorm:"not null VARCHAR(40) unique(dtl)" json:"productCode,omitempty" xml:"productCode"`
	ProductName  string  `xorm:"VARCHAR(200)" json:"productName,omitempty" xml:"productName"`
	FromLocation string  `xorm:"comment('sqlacc item basis') VARCHAR(200)" json:"fromLocation,omitempty" xml:"fromLocation"`
	ToLocation   string  `xorm:"comment('sqlacc item basis') VARCHAR(200)" json:"toLocation,omitempty" xml:"toLocation"`
	Quantity     int     `xorm:"INT" json:"quantity,omitempty" xml:"quantity"`
	UnitUom      string  `xorm:"VARCHAR(40) unique(dtl)" json:"unitUom,omitempty" xml:"unitUom"`
	UnitPrice    float64 `xorm:"comment('unit cost') DOUBLE" json:"unitPrice,omitempty" xml:"unitPrice"`
	SubTotal     float64 `xorm:"DOUBLE" json:"subTotal,omitempty" xml:"subTotal"`
	AgentRemark  string  `xorm:"VARCHAR(500)" json:"agentRemark,omitempty" xml:"agentRemark"`
	CancelStatus int     `xorm:"default 0 INT" json:"cancelStatus,omitempty" xml:"cancelStatus"`
}

func (m *CmsStockTransferDtl) TableName() string {
	return "cms_stock_transfer_dtl"
}
