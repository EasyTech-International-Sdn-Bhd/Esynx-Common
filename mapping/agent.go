package mapping

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-common/forms"
)

func MapToAgent(d forms.CmsLoginForm) entities.CmsLogin {
	return entities.CmsLogin{
		AgentCode:   d.AgentCode,
		Login:       d.Login,
		Password:    d.Password,
		Name:        d.Name,
		Email:       d.Email,
		ContactNo:   d.ContactNo,
		DeviceToken: d.DeviceToken,
		RoleId:      d.RoleId,
		Remark:      d.Remark,
		LoginStatus: d.LoginStatus,
		DocSuffix:   d.DocSuffix,
		ProjNo:      d.ProjNo,
		Company:     d.Company,
	}
}
