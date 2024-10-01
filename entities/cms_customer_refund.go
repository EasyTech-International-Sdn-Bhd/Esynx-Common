package entities

import (
	"time"
)

type CmsCustomerRefund struct {
	CfId             uint64    `xorm:"pk autoincr UNSIGNED BIGINT" json:"cfId,omitempty" xml:"cfId"`
	CfCode           string    `xorm:"index unique 'cfCode' VARCHAR(20)" json:"cfCode,omitempty" xml:"cfCode"`
	CustCode         string    `xorm:"index VARCHAR(20)" json:"custCode,omitempty" xml:"custCode"`
	CfDate           time.Time `xorm:"TIMESTAMP" json:"cfDate,omitempty" xml:"cfDate"`
	CfAmount         float64   `xorm:"DOUBLE" json:"cfAmount,omitempty" xml:"cfAmount"`
	Cancelled        string    `xorm:"CHAR(1)" json:"cancelled,omitempty" xml:"cancelled"`
	CfKnockoffAmount float64   `xorm:"DOUBLE" json:"cfKnockoffAmount,omitempty" xml:"cfKnockoffAmount"`
	AgentCode        string    `xorm:"VARCHAR(20)" json:"agentCode,omitempty" xml:"agentCode"`
	RefNo            string    `xorm:"VARCHAR(50)" json:"refNo,omitempty" xml:"refNo"`
	UpdatedAt        time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsCustomerRefund) TableName() string {
	return "cms_customer_refund"
}
