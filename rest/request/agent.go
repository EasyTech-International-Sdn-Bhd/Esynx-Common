package request

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/forms"
)

type AgentCreateForm struct {
	Reference string             `json:"reference" binding:"required,min=3,max=100"`
	Data      forms.CmsLoginForm `json:"data"`
}

type AgentBulkCreateForm struct {
	Data []AgentCreateForm `json:"data"`
}

type AgentBulkUpdateForm struct {
	Data []AgentCreateForm `json:"data"`
}

type AgentDeleteForm struct {
	AgentCode string `json:"agentCode" binding:"required,min=3,max=100"`
}

type AgentBulkDeleteForm struct {
	Data []AgentDeleteForm `json:"data"`
}
