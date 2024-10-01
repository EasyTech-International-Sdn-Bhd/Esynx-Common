package entities

import (
	"time"
)

type CmsStockTransfer struct {
	Id             uint64    `xorm:"pk autoincr unique(stCode) UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	StCode         string    `xorm:"not null unique(stCode) VARCHAR(200)" json:"stCode,omitempty" xml:"stCode"`
	StDate         time.Time `xorm:"not null DATETIME" json:"stDate,omitempty" xml:"stDate"`
	CustCode       string    `xorm:"comment('can be empty') VARCHAR(200)" json:"custCode,omitempty" xml:"custCode"`
	FromLocation   string    `xorm:"VARCHAR(200)" json:"fromLocation,omitempty" xml:"fromLocation"`
	ToLocation     string    `xorm:"VARCHAR(200)" json:"toLocation,omitempty" xml:"toLocation"`
	Total          float64   `xorm:"DOUBLE" json:"total,omitempty" xml:"total"`
	Note           string    `xorm:"VARCHAR(200)" json:"note,omitempty" xml:"note"`
	StStatus       int       `xorm:"comment('0 - in app; 1 - confirm; 2 - transferred to ATC') INT" json:"stStatus,omitempty" xml:"stStatus"`
	CancelStatus   int       `xorm:"default 0 INT" json:"cancelStatus,omitempty" xml:"cancelStatus"`
	AgentCode      string    `xorm:"VARCHAR(20)" json:"agentCode,omitempty" xml:"agentCode"`
	StFault        int       `xorm:"default 0 INT" json:"stFault,omitempty" xml:"stFault"`
	StFaultMessage string    `xorm:"VARCHAR(200)" json:"stFaultMessage,omitempty" xml:"stFaultMessage"`
}

func (m *CmsStockTransfer) TableName() string {
	return "cms_stock_transfer"
}
