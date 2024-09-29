package forms

import "github.com/easytech-international-sdn-bhd/esynx-common/forms"

type DebitNoteFormEDA struct {
	Head DefaultForm               `json:"head" binding:"required"`
	Data []forms.CmsCreditNoteForm `json:"data" binding:"required"`
}

type DebitNoteDetailsFormEDA struct {
	Head DefaultForm                      `json:"head" binding:"required"`
	Data []forms.CmsCreditNoteDetailsForm `json:"data" binding:"required"`
}
