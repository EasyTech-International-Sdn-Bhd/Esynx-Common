package entities

import "time"

type AuditLog struct {
	OperationType string    `xorm:"ENUM('INSERT', 'UPDATE', 'ERROR') not null"`
	RecordTable   string    `xorm:"VARCHAR(100)"`
	RecordId      string    `xorm:"VARCHAR(80)"`
	RecordBody    string    `xorm:"JSON"`
	UserCode      string    `xorm:"VARCHAR(50) not null"`
	AppName       string    `xorm:"VARCHAR(20) not null"`
	OperationTime time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (m *AuditLog) TableName() string {
	return "audit_log"
}
