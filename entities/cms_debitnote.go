package entities

import (
	"time"
)

type CmsDebitnote struct {
	DnId              uint64    `xorm:"pk autoincr UNSIGNED BIGINT" json:"dnId,omitempty" xml:"dnId"`
	DnCode            string    `xorm:"index VARCHAR(40)" json:"dnCode,omitempty" xml:"dnCode"`
	CustCode          string    `xorm:"index VARCHAR(40)" json:"custCode,omitempty" xml:"custCode"`
	DnDate            time.Time `xorm:"TIMESTAMP" json:"dnDate,omitempty" xml:"dnDate"`
	DnAmount          float64   `xorm:"DOUBLE" json:"dnAmount,omitempty" xml:"dnAmount"`
	OutstandingAmount float64   `xorm:"DOUBLE" json:"outstandingAmount,omitempty" xml:"outstandingAmount"`
	AgentCode         string    `xorm:"VARCHAR(20)" json:"agentCode,omitempty" xml:"agentCode"`
	Cancelled         string    `xorm:"CHAR(1)" json:"cancelled,omitempty" xml:"cancelled"`
	Approved          int       `xorm:"default 0 INT" json:"approved,omitempty" xml:"approved"`
	Approver          string    `xorm:"VARCHAR(200)" json:"approver,omitempty" xml:"approver"`
	ApprovedAt        time.Time `xorm:"DATETIME" json:"approvedAt,omitempty" xml:"approvedAt"`
	FromDoc           string    `xorm:"default 'SL' ENUM('AR','SL')" json:"fromDoc,omitempty" xml:"fromDoc"`
	UpdatedAt         time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
	RefNo             string    `xorm:"VARCHAR(80) index unique" json:"refNo,omitempty" xml:"refNo"`
	DnUdf             string    `xorm:"JSON" json:"dnUdf,omitempty" xml:"dnUdf"`
}

func (m *CmsDebitnote) TableName() string {
	return "cms_debitnote"
}

func (m *CmsDebitnote) BeforeInsert() {
	m.BeforeUpdate()
}

func (m *CmsDebitnote) BeforeUpdate() {
	m.UpdatedAt = time.Now()
}
