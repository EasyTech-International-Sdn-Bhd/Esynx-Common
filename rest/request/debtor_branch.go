package request

import "github.com/easytech-international-sdn-bhd/esynx-common/forms"

type DebtorBranchCreateOrUpdateForm struct {
	Reference string                      `json:"reference" binding:"required,min=3,max=100"`
	Data      forms.CmsCustomerBranchForm `json:"data"`
}

type DebtorBranchCreateOrUpdateBulkForm struct {
	Data []DebtorBranchCreateOrUpdateForm `json:"data"`
}

type DebtorBranchDeleteForm struct {
	Reference string `json:"reference" binding:"required,min=3,max=100"`
}

type DebtorBranchDeleteBulkForm struct {
	Data []DebtorBranchDeleteForm `json:"data"`
}
