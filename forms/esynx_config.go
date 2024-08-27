package forms

type EsynxConfigForm struct {
	Name      string `validate:"trim" xorm:"unique(unq) VARCHAR(20) not null" json:"name" xml:"name"`
	ServiceId string `validate:"trim" xorm:"unique(unq) VARCHAR(20) not null" json:"serviceId" xml:"serviceId"`
	Config    string `validate:"trim" xorm:"JSON" json:"config" xml:"config"`
}
