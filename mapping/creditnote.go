package mapping

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-common/forms"
)

func MapToCreditNote(dt forms.CmsCreditNoteForm) entities.CmsCreditnote {
	return entities.CmsCreditnote{
		CnCode:           dt.CnCode,
		CustCode:         dt.CustCode,
		CnDate:           dt.CnDate,
		CnAmount:         dt.CnAmount,
		CnKnockoffAmount: dt.CnKnockoffAmount,
		Approved:         dt.Approved,
		Approver:         dt.Approver,
		ApprovedAt:       dt.ApprovedAt,
		CnUdf:            dt.CnUdf,
		RefNo:            dt.RefNo,
		FromDoc:          dt.FromDoc,
		Cancelled:        dt.Cancelled,
		AgentCode:        dt.Agent,
	}
}

func MapToCreditNoteDetails(dt forms.CmsCreditNoteDetailsForm) entities.CmsCreditnoteDetails {
	return entities.CmsCreditnoteDetails{
		CnCode:       dt.CnCode,
		ItemCode:     dt.ItemCode,
		ItemName:     dt.ItemName,
		ItemPrice:    dt.ItemPrice,
		Quantity:     dt.Quantity,
		TotalPrice:   dt.TotalPrice,
		Uom:          dt.Uom,
		Discount:     dt.Discount,
		ActiveStatus: dt.ActiveStatus,
		SequenceNo:   dt.SequenceNo,
		CnDtlUdf:     dt.CnDtlUdf,
		RefNo:        dt.RefNo,
	}
}
