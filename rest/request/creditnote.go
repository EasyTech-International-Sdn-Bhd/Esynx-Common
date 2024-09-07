package request

import "github.com/easytech-international-sdn-bhd/esynx-common/forms"

type CreditNoteCreateForm struct {
	Reference string                  `json:"reference" binding:"required,min=3,max=100"`
	DocType   string                  `json:"docType" binding:"required,min=2"`
	Data      forms.CmsCreditNoteForm `json:"data" binding:"required"`
}

type CreditNoteBulkCreateForm struct {
	Data []CreditNoteCreateForm `json:"data" binding:"required,v-nonempty"`
}

type CreditNoteUpdateForm struct {
	Reference string                  `json:"reference" binding:"required,min=3,max=100"`
	Data      forms.CmsCreditNoteForm `json:"data" binding:"required"`
}

type CreditNoteBulkUpdateForm struct {
	Data []CreditNoteUpdateForm `json:"data" binding:"required,v-nonempty"`
}

type CreditNoteDeleteForm struct {
	CreditNoteCode string `json:"cnCode" binding:"required,min=3,max=100"`
}

type CreditNoteBulkDeleteForm struct {
	Data []CreditNoteDeleteForm `json:"cnCodes" binding:"required,v-nonempty"`
}

type CreditNoteDetailsCreateForm struct {
	Reference string                         `json:"reference" binding:"required,min=3,max=100"`
	Data      forms.CmsCreditNoteDetailsForm `json:"data" binding:"required"`
}

type CreditNoteDetailsBulkCreateForm struct {
	Data []CreditNoteDetailsCreateForm `json:"data" binding:"required,v-nonempty"`
}

type CreditNoteDetailsUpdateForm struct {
	Reference string                         `json:"reference" binding:"required,min=3,max=100"`
	Data      forms.CmsCreditNoteDetailsForm `json:"data" binding:"required"`
}

type CreditNoteDetailsBulkUpdateForm struct {
	Data []CreditNoteDetailsUpdateForm `json:"data" binding:"required,v-nonempty"`
}

type CreditNoteDetailsDeleteForm struct {
	Reference string `json:"id" binding:"required,min=3,max=100"`
}

type CreditNoteDetailsBulkDeleteForm struct {
	Data []CreditNoteDetailsDeleteForm `json:"ids" binding:"required,v-nonempty"`
}
