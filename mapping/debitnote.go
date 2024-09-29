package mapping

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-common/forms"
)

func MapToDebitNote(dt forms.CmsDebitNoteForm) entities.CmsDebitnote {
	return entities.CmsDebitnote{
		DnCode:            dt.DnCode,
		CustCode:          dt.CustCode,
		DnDate:            dt.DnDate,
		DnAmount:          dt.DnAmount,
		OutstandingAmount: dt.OutstandingAmount,
		Approved:          dt.Approved,
		Approver:          dt.Approver,
		ApprovedAt:        dt.ApprovedAt,
		DnUdf:             dt.DnUdf,
		RefNo:             dt.RefNo,
		FromDoc:           dt.FromDoc,
		Cancelled:         dt.Cancelled,
		AgentCode:         dt.Agent,
	}
}

func MapToDebitNoteDetails(dt forms.CmsDebitNoteDetailsForm) entities.CmsDebitnoteDetails {
	return entities.CmsDebitnoteDetails{
		DnCode:       dt.DnCode,
		ItemCode:     dt.ItemCode,
		ItemName:     dt.ItemName,
		ItemPrice:    dt.ItemPrice,
		Quantity:     dt.Quantity,
		TotalPrice:   dt.TotalPrice,
		Uom:          dt.Uom,
		Discount:     dt.Discount,
		ActiveStatus: dt.ActiveStatus,
		SequenceNo:   dt.SequenceNo,
		DnDtlUdf:     dt.DnDtlUdf,
		RefNo:        dt.RefNo,
	}
}
