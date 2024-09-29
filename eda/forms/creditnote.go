package forms

import "github.com/easytech-international-sdn-bhd/esynx-common/forms"

type CreditNoteFormEDA struct {
	Head DefaultForm               `json:"head" binding:"required"`
	Data []forms.CmsCreditNoteForm `json:"data" binding:"required"`
}

type CreditNoteDetailsFormEDA struct {
	Head DefaultForm                      `json:"head" binding:"required"`
	Data []forms.CmsCreditNoteDetailsForm `json:"data" binding:"required"`
}
