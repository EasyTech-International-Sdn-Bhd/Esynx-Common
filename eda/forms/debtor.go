package forms

import "github.com/easytech-international-sdn-bhd/esynx-common/forms"

type DebtorFormEDA struct {
	Head DefaultForm             `json:"head" binding:"required"`
	Data []forms.CmsCustomerForm `json:"data" binding:"required"`
}