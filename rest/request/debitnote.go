package request

import "github.com/easytech-international-sdn-bhd/esynx-common/forms"

type DebitNoteCreateForm struct {
	Reference string                 `json:"reference" binding:"required,min=3,max=100"`
	DocType   string                 `json:"docType" binding:"required,min=2"`
	Data      forms.CmsDebitNoteForm `json:"data"`
}

type DebitNoteBulkCreateForm struct {
	Data []DebitNoteCreateForm `json:"data"`
}

type DebitNoteUpdateForm struct {
	Reference string                 `json:"reference" binding:"required,min=3,max=100"`
	Data      forms.CmsDebitNoteForm `json:"data"`
}

type DebitNoteBulkUpdateForm struct {
	Data []DebitNoteUpdateForm `json:"data"`
}

type DebitNoteDeleteForm struct {
	DebitNoteCode string `json:"dnCode" binding:"required,min=3,max=100"`
}

type DebitNoteBulkDeleteForm struct {
	Data []DebitNoteDeleteForm `json:"dnCodes"`
}

type DebitNoteDetailsCreateForm struct {
	Reference string                        `json:"reference" binding:"required,min=3,max=100"`
	Data      forms.CmsDebitNoteDetailsForm `json:"data"`
}

type DebitNoteDetailsBulkCreateForm struct {
	Data []DebitNoteDetailsCreateForm `json:"data"`
}

type DebitNoteDetailsUpdateForm struct {
	Reference string                        `json:"reference" binding:"required,min=3,max=100"`
	Data      forms.CmsDebitNoteDetailsForm `json:"data"`
}

type DebitNoteDetailsBulkUpdateForm struct {
	Data []DebitNoteDetailsUpdateForm `json:"data"`
}

type DebitNoteDetailsDeleteForm struct {
	Reference string `json:"id" binding:"required,min=3,max=100"`
}

type DebitNoteDetailsBulkDeleteForm struct {
	Data []DebitNoteDetailsDeleteForm `json:"ids"`
}
