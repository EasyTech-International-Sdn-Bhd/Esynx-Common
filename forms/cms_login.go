package forms

type CmsLoginForm struct {
	StaffCode   string `xorm:"unique VARCHAR(20)" json:"staffCode,omitempty" xml:"staffCode" validate:"trim"`
	Login       string `xorm:"VARCHAR(30)" json:"login,omitempty" xml:"login" validate:"trim"`
	Password    string `xorm:"VARCHAR(30)" json:"password,omitempty" xml:"password" validate:"trim"`
	Name        string `xorm:"VARCHAR(200)" json:"name,omitempty" xml:"name" validate:"trim"`
	Email       string `xorm:"VARCHAR(250)" json:"email,omitempty" xml:"email" validate:"trim"`
	ContactNo   string `xorm:"VARCHAR(100)" json:"contactNo,omitempty" xml:"contactNo" validate:"trim"`
	DeviceToken string `xorm:"VARCHAR(100)" json:"deviceToken,omitempty" xml:"deviceToken" validate:"trim"`
	Remark      string `xorm:"BLOB" json:"remark,omitempty" xml:"remark" validate:"trim"`
	DocSuffix   string `xorm:"default 'S' VARCHAR(10)" json:"docSuffix,omitempty" xml:"docSuffix" validate:"trim"`
	ProjNo      string `xorm:"VARCHAR(100)" json:"projNo,omitempty" xml:"projNo" validate:"trim"`
	SessionId   string `xorm:"default '' VARCHAR(100)" json:"sessionId,omitempty" xml:"sessionId" validate:"trim"`
	Company     string `xorm:"default '' VARCHAR(50)" json:"company,omitempty" xml:"company" validate:"trim"`
	RoleId      int    `xorm:"default 2 comment('Officer, Salesperson, Admin') INT" json:"roleId,omitempty" xml:"roleId"`
	LoginStatus int    `xorm:"default 1 comment('1=active, 0=inactive , please check the disable salesperson is not allowed to send in order.') INT" json:"loginStatus,omitempty" xml:"loginStatus"`
}
