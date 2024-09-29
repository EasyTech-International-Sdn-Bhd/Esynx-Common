package mapping

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-common/forms"
)

func MapToInvoice(d forms.CmsInvoiceForm) entities.CmsInvoice {
	return entities.CmsInvoice{
		InvoiceCode:       d.InvoiceCode,
		CustCode:          d.CustCode,
		InvoiceDate:       d.InvoiceDate,
		InvoiceDueDate:    d.InvoiceDueDate,
		InvoiceAmount:     d.InvoiceAmount,
		OutstandingAmount: d.OutstandingAmount,
		Approved:          d.Approved,
		Approver:          d.Approver,
		ApprovedAt:        d.ApprovedAt,
		InvUdf:            d.InvUdf,
		RefNo:             d.RefNo,
		DocType:           d.DocType,
		FromDoc:           d.FromDoc,
		Cancelled:         d.Cancelled,
		Termcode:          d.Termcode,
		AgentCode:         d.Agent,
	}
}

func MapToInvoiceDetails(d forms.CmsInvoiceDetailsForm) entities.CmsInvoiceDetails {
	return entities.CmsInvoiceDetails{
		InvoiceCode:  d.InvoiceCode,
		ItemCode:     d.ItemCode,
		ItemName:     d.ItemName,
		ItemPrice:    d.ItemPrice,
		Quantity:     d.Quantity,
		TotalPrice:   d.TotalPrice,
		Uom:          d.Uom,
		Discount:     d.Discount,
		ActiveStatus: d.ActiveStatus,
		SequenceNo:   d.SequenceNo,
		InvDtlUdf:    d.InvDtlUdf,
		RefNo:        d.RefNo,
	}
}
