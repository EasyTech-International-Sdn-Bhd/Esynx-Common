package forms

import "github.com/easytech-international-sdn-bhd/esynx-common/forms"

type AgentFormEDA struct {
	Head EdaHeader            `json:"head" binding:"required"`
	Data []forms.CmsLoginForm `json:"data" binding:"required"`
}
