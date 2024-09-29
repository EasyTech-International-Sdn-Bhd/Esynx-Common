package mapping

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-common/forms"
)

func MapToDebtor(dt forms.CmsCustomerForm) entities.CmsCustomer {
	return entities.CmsCustomer{
		CreatedDate:        dt.CreatedDate,
		CustCode:           dt.CustCode,
		CustCompanyName:    dt.CustCompanyName,
		CustInchargePerson: dt.CustInchargePerson,
		CustRemark:         dt.CustRemark,
		CustReference:      dt.CustReference,
		CustEmail:          dt.CustEmail,
		CustTel:            dt.CustTel,
		CustFax:            dt.CustFax,
		BillingAddress1:    dt.BillingAddress1,
		BillingAddress2:    dt.BillingAddress2,
		BillingAddress3:    dt.BillingAddress3,
		BillingAddress4:    dt.BillingAddress4,
		BillingCity:        dt.BillingCity,
		BillingState:       dt.BillingState,
		BillingZipcode:     dt.BillingZipcode,
		BillingCountry:     dt.BillingCountry,
		ShippingAddress1:   dt.ShippingAddress1,
		ShippingAddress2:   dt.ShippingAddress2,
		ShippingAddress3:   dt.ShippingAddress3,
		ShippingAddress4:   dt.ShippingAddress4,
		ShippingCity:       dt.ShippingCity,
		ShippingState:      dt.ShippingState,
		ShippingZipcode:    dt.ShippingZipcode,
		ShippingCountry:    dt.ShippingCountry,
		Termcode:           dt.Termcode,
		SellingPriceType:   dt.SellingPriceType,
		CustomerZone:       dt.CustomerZone,
		CustomerStatus:     dt.CustomerStatus,
		CreditLimit:        dt.CreditLimit,
		CurrentBalance:     dt.CurrentBalance,
		Currency:           dt.Currency,
		CurrencyRate:       dt.CurrencyRate,
		Latitude:           dt.Latitude,
		Longitude:          dt.Longitude,
		CustUdf:            dt.CustUdf,
		Company:            dt.Company,
	}
}
