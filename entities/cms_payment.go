package entities

import (
	"time"
)

type CmsPayment struct {
	PaymentId                   string    `xorm:"unique(id) VARCHAR(40) index" json:"paymentId,omitempty" xml:"paymentId"`
	PaymentDate                 time.Time `xorm:"DATETIME" json:"paymentDate,omitempty" xml:"paymentDate"`
	PaymentTransferReceivedDate time.Time `xorm:"DATETIME" json:"paymentTransferReceivedDate,omitempty" xml:"paymentTransferReceivedDate"`
	CustCode                    string    `xorm:"VARCHAR(50) index" json:"custCode,omitempty" xml:"custCode"`
	Description                 string    `xorm:"VARCHAR(500)" json:"description,omitempty" xml:"description"`
	PaymentAmount               float64   `xorm:"DOUBLE" json:"paymentAmount,omitempty" xml:"paymentAmount"`
	PaymentStatus               int       `xorm:"comment('0=in_ipad,1=in_backoffice,2=in_Accounting') INT" json:"paymentStatus,omitempty" xml:"paymentStatus"`
	CancelStatus                int       `xorm:"INT" json:"cancelStatus,omitempty" xml:"cancelStatus"`
	PaymentReference            string    `xorm:"VARCHAR(50)" json:"paymentReference,omitempty" xml:"paymentReference"`
	AgentCode                   string    `xorm:"VARCHAR(20) index" json:"agentCode,omitempty" xml:"agentCode"`
	AgentPaymentRemark          string    `xorm:"VARCHAR(500)" json:"agentPaymentRemark,omitempty" xml:"agentPaymentRemark"`
	LastPrint                   time.Time `xorm:"DATETIME" json:"lastPrint,omitempty" xml:"lastPrint"`
	Checked                     int       `xorm:"default 0 INT" json:"checked,omitempty" xml:"checked"`
	PaymentFault                int       `xorm:"default 0 INT" json:"paymentFault,omitempty" xml:"paymentFault"`
	PaymentFaultMessage         string    `xorm:"VARCHAR(200)" json:"paymentFaultMessage,omitempty" xml:"paymentFaultMessage"`
	PaymentStatusLastUpdateDate time.Time `xorm:"DATETIME" json:"paymentStatusLastUpdateDate,omitempty" xml:"paymentStatusLastUpdateDate"`
	PaymentStatusLastUpdateBy   string    `xorm:"VARCHAR(50)" json:"paymentStatusLastUpdateBy,omitempty" xml:"paymentStatusLastUpdateBy"`
	KnockoffInv                 string    `xorm:"JSON" json:"knockoffInv,omitempty" xml:"knockoffInv"`
	DocType                     string    `xorm:"default 'payment' VARCHAR(10)" json:"docType,omitempty" xml:"docType"`
	DocId                       string    `xorm:"VARCHAR(20)" json:"docId,omitempty" xml:"docId"`
}

func (m *CmsPayment) TableName() string {
	return "cms_payment"
}
