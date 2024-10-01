package entities

type CmsModule struct {
	ModuleId uint64 `xorm:"pk autoincr unique(id) UNSIGNED BIGINT" json:"moduleId,omitempty" xml:"moduleId"`
	Name     string `xorm:"not null unique(id) VARCHAR(100)" json:"name,omitempty" xml:"name"`
	Value    []byte `xorm:"BLOB" json:"value,omitempty" xml:"value"`
}

func (m *CmsModule) TableName() string {
	return "cms_module"
}
