package forms

import "github.com/easytech-international-sdn-bhd/esynx-common/forms"

type DebtorBranchFormEDA struct {
	Head EdaHeader                     `json:"head" binding:"required"`
	Data []forms.CmsCustomerBranchForm `json:"data" binding:"required"`
}
