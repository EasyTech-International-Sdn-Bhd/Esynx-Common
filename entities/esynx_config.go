package entities

import (
	"time"
)

type EsynxConfig struct {
	Id        int64     `xorm:"pk autoincr" json:"id" xml:"id"`
	Name      string    `xorm:"unique(unq) VARCHAR(20) not null" json:"name" xml:"name"`
	ServiceId string    `xorm:"unique(unq) VARCHAR(20) not null" json:"serviceId" xml:"serviceId"`
	Config    string    `xorm:"JSON" json:"config" xml:"config"`
	CreatedAt time.Time `xorm:"TIMESTAMP" json:"createdAt" xml:"createdAt"`
	UpdatedAt time.Time `xorm:"TIMESTAMP" json:"updatedAt" xml:"updatedAt"`
}

func (m *EsynxConfig) TableName() string {
	return "esynx_config"
}

func (m *EsynxConfig) BeforeInsert() {
	m.CreatedAt = time.Now()
	m.BeforeUpdate()
}

func (m *EsynxConfig) BeforeUpdate() {
	m.UpdatedAt = time.Now()
}
