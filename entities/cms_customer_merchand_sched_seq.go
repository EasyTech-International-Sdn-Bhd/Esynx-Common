package entities

import (
	"time"
)

type CmsCustomerMerchandSchedSeq struct {
	Id           uint64    `xorm:"pk autoincr unique(appmntId) UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	AppmntDtlId  int       `xorm:"not null unique(appmntId) INT" json:"appmntDtlId,omitempty" xml:"appmntDtlId"`
	AgentCode    string    `xorm:"VARCHAR(20)" json:"agentCode,omitempty" xml:"agentCode"`
	Sequence     int       `xorm:"not null INT" json:"sequence,omitempty" xml:"sequence"`
	ActiveStatus int       `xorm:"default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	UpdatedAt    time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsCustomerMerchandSchedSeq) TableName() string {
	return "cms_customer_merchand_sched_seq"
}
