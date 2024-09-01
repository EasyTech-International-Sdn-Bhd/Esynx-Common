package request

import "github.com/easytech-international-sdn-bhd/esynx-common/forms"

type DebtorCreateOrUpdateForm struct {
	Reference string                `json:"reference" binding:"required,min=3,max=100"`
	Data      forms.CmsCustomerForm `json:"data"`
}

type DebtorCreateOrUpdateBulkForm struct {
	Data []DebtorCreateOrUpdateForm `json:"data"`
}

type DebtorDeleteForm struct {
	Reference string `json:"reference" binding:"required,min=3,max=100"`
}

type DebtorDeleteBulkForm struct {
	Data []DebtorDeleteForm `json:"data"`
}

type DebtorAssignAgentForm struct {
	CustCode  string `json:"custCode" binding:"required,min=3,max=100"`
	AgentCode string `json:"agentCode" binding:"required,min=3,max=100"`
}
