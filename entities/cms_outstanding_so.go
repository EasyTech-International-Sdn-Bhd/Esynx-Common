package entities

import (
	"time"
)

type CmsOutstandingSo struct {
	Id           uint64    `xorm:"pk autoincr unique(docNo) UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	DocNo        string    `xorm:"not null unique(docNo) index VARCHAR(20)" json:"docNo,omitempty" xml:"docNo"`
	DocKey       string    `xorm:"not null VARCHAR(30)" json:"docKey,omitempty" xml:"docKey"`
	ProductCode  string    `xorm:"not null default '' index unique(docNo) VARCHAR(50)" json:"productCode,omitempty" xml:"productCode"`
	UnitUom      string    `xorm:"unique(docNo) VARCHAR(50)" json:"unitUom,omitempty" xml:"unitUom"`
	OriQty       float64   `xorm:"not null unique(docNo) DOUBLE" json:"oriQty,omitempty" xml:"oriQty"`
	OutQty       float64   `xorm:"not null DOUBLE" json:"outQty,omitempty" xml:"outQty"`
	TransQty     float64   `xorm:"not null DOUBLE" json:"transQty,omitempty" xml:"transQty"`
	DocDate      time.Time `xorm:"not null DATETIME" json:"docDate,omitempty" xml:"docDate"`
	AgentCode    string    `xorm:"VARCHAR(20) index" json:"agentCode,omitempty" xml:"agentCode"`
	CustCode     string    `xorm:"not null VARCHAR(50) index" json:"custCode,omitempty" xml:"custCode"`
	BranchCode   string    `xorm:"VARCHAR(100)" json:"branchCode,omitempty" xml:"branchCode"`
	UpdatedAt    time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
	ActiveStatus int       `xorm:"default 0 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	RefNo        string    `xorm:"VARCHAR(100)" json:"refNo,omitempty" xml:"refNo"`
}

func (m *CmsOutstandingSo) TableName() string {
	return "cms_outstanding_so"
}
