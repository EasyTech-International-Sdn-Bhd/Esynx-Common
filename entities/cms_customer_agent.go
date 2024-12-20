package entities

import (
	"time"
)

type CmsCustomerAgent struct {
	Id           uint64    `xorm:"pk autoincr UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	AgentCode    string    `xorm:"unique(unique_customer) VARCHAR(20)" json:"agentCode,omitempty" xml:"agentCode"`
	CustCode     string    `xorm:"unique(unique_customer) VARCHAR(40)" json:"custCode,omitempty" xml:"custCode"`
	ActiveStatus int       `xorm:"default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	UpdatedAt    time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsCustomerAgent) TableName() string {
	return "cms_customer_agent"
}

func (m *CmsCustomerAgent) Validate() {

}

func (m *CmsCustomerAgent) BeforeInsert() {
	m.BeforeUpdate()
}

func (m *CmsCustomerAgent) BeforeUpdate() {
	m.UpdatedAt = time.Now()
}
