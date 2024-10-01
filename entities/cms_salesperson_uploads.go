package entities

import (
	"time"
)

type CmsUserUploads struct {
	Id         uint64    `xorm:"pk autoincr unique(unx) UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	Image      string    `xorm:"not null unique(unx) VARCHAR(200)" json:"image,omitempty" xml:"image"`
	TypeName   string    `xorm:"not null comment('module name from cms_module_camera') VARCHAR(200)" json:"typeName,omitempty" xml:"typeName"`
	Remark     string    `xorm:"VARCHAR(200)" json:"uploadRemark,omitempty" xml:"uploadRemark"`
	Status     int       `xorm:"not null default 1 comment('1 - not deleted 0 - deleted') INT" json:"status,omitempty" xml:"status"`
	AgentCode  string    `xorm:"not null INT" json:"agentCode,omitempty" xml:"agentCode"`
	Location   string    `xorm:"VARCHAR(200)" json:"location,omitempty" xml:"location"`
	BindId     string    `xorm:"not null comment('id - from mobile') VARCHAR(200)" json:"bindId,omitempty" xml:"bindId"`
	BindType   string    `xorm:"not null comment('SO/CO') VARCHAR(200)" json:"bindType,omitempty" xml:"bindType"`
	UploadDate time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"uploadDate,omitempty" xml:"uploadDate"`
	TakenDate  time.Time `xorm:"DATETIME" json:"takenDate,omitempty" xml:"takenDate"`
}

func (m *CmsUserUploads) TableName() string {
	return "cms_user_uploads"
}
