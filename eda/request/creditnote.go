package request

import "github.com/easytech-international-sdn-bhd/esynx-common/forms"

type CreditNoteFormEDA struct {
	Head EdaHeader                 `json:"head" binding:"required"`
	Data []forms.CmsCreditNoteForm `json:"data" binding:"required"`
}

type CreditNoteDetailsFormEDA struct {
	Head EdaHeader                        `json:"head" binding:"required"`
	Data []forms.CmsCreditNoteDetailsForm `json:"data" binding:"required"`
}
