package request

import "github.com/easytech-international-sdn-bhd/esynx-common/forms"

type DebitNoteCreateForm struct {
	Reference string                 `json:"reference" binding:"required,min=3,max=100"`
	Data      forms.CmsDebitNoteForm `json:"data" binding:"required"`
}

type DebitNoteBulkCreateForm struct {
	Data []DebitNoteCreateForm `json:"data" binding:"required,v-nonempty"`
}

type DebitNoteUpdateForm struct {
	Reference string                 `json:"reference" binding:"required,min=3,max=100"`
	Data      forms.CmsDebitNoteForm `json:"data" binding:"required"`
}

type DebitNoteBulkUpdateForm struct {
	Data []DebitNoteUpdateForm `json:"data" binding:"required,v-nonempty"`
}

type DebitNoteDeleteForm struct {
	DebitNoteCode string `json:"dnCode" binding:"required,min=3,max=100"`
}

type DebitNoteBulkDeleteForm struct {
	Data []DebitNoteDeleteForm `json:"dnCodes" binding:"required,v-nonempty"`
}

type DebitNoteDetailsCreateForm struct {
	Reference string                        `json:"reference" binding:"required,min=3,max=100"`
	Data      forms.CmsDebitNoteDetailsForm `json:"data" binding:"required"`
}

type DebitNoteDetailsBulkCreateForm struct {
	Data []DebitNoteDetailsCreateForm `json:"data" binding:"required,v-nonempty"`
}

type DebitNoteDetailsUpdateForm struct {
	Reference string                        `json:"reference" binding:"required,min=3,max=100"`
	Data      forms.CmsDebitNoteDetailsForm `json:"data" binding:"required"`
}

type DebitNoteDetailsBulkUpdateForm struct {
	Data []DebitNoteDetailsUpdateForm `json:"data" binding:"required,v-nonempty"`
}

type DebitNoteDetailsDeleteForm struct {
	Reference string `json:"id" binding:"required,min=3,max=100"`
}

type DebitNoteDetailsBulkDeleteForm struct {
	Data []DebitNoteDetailsDeleteForm `json:"ids" binding:"required,v-nonempty"`
}
