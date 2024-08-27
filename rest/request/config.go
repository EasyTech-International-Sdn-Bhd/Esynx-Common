package request

import "github.com/easytech-international-sdn-bhd/esynx-common/forms"

type ConfigForm struct {
	Reference string                `json:"reference" binding:"required"`
	Data      forms.EsynxConfigForm `json:"data"`
}

type ConfigDeleteForm struct {
	ServiceId string `json:"serviceId" binding:"required"`
}
