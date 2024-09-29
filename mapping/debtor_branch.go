package mapping

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-common/forms"
)

func MapToDebtorBranch(dt forms.CmsCustomerBranchForm) entities.CmsCustomerBranch {
	return entities.CmsCustomerBranch{
		AgentCode:        dt.Agent,
		CustCode:         dt.CustCode,
		BranchCode:       dt.BranchCode,
		BranchName:       dt.BranchName,
		BranchAttn:       dt.BranchAttn,
		BranchPhone:      dt.BranchPhone,
		BranchFax:        dt.BranchFax,
		BillingAddress1:  dt.BillingAddress1,
		BillingAddress2:  dt.BillingAddress2,
		BillingAddress3:  dt.BillingAddress3,
		BillingAddress4:  dt.BillingAddress4,
		BillingState:     dt.BillingState,
		BillingPostcode:  dt.BillingPostcode,
		BillingCountry:   dt.BillingCountry,
		ShippingAddress1: dt.ShippingAddress1,
		ShippingAddress2: dt.ShippingAddress2,
		ShippingAddress3: dt.ShippingAddress3,
		ShippingAddress4: dt.ShippingAddress4,
		ShippingState:    dt.ShippingState,
		ShippingPostcode: dt.ShippingPostcode,
		ShippingCountry:  dt.ShippingCountry,
		BranchArea:       dt.BranchArea,
		BranchRemark:     dt.BranchRemark,
		BranchActive:     dt.BranchActive,
		Company:          dt.Company,
		BranchUdf:        dt.BranchUdf,
	}
}
