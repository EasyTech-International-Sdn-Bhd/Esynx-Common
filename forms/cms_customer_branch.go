package forms

type CmsCustomerBranchForm struct {
	Agent            string `xorm:"not null index VARCHAR(20)" json:"agent,omitempty" xml:"agent"`
	CustCode         string `xorm:"not null index unique(unique_branch) VARCHAR(40)" json:"custCode,omitempty" xml:"custCode"`
	BranchCode       string `xorm:"not null index unique(unique_branch) VARCHAR(40)" json:"branchCode,omitempty" xml:"branchCode"`
	BranchName       string `xorm:"not null VARCHAR(200)" json:"branchName,omitempty" xml:"branchName"`
	BranchAttn       string `xorm:"not null VARCHAR(100)" json:"branchAttn,omitempty" xml:"branchAttn"`
	BranchPhone      string `xorm:"not null VARCHAR(100)" json:"branchPhone,omitempty" xml:"branchPhone"`
	BranchFax        string `xorm:"not null VARCHAR(100)" json:"branchFax,omitempty" xml:"branchFax"`
	BillingAddress1  string `xorm:"not null VARCHAR(200)" json:"billingAddress1,omitempty" xml:"billingAddress1"`
	BillingAddress2  string `xorm:"not null VARCHAR(200)" json:"billingAddress2,omitempty" xml:"billingAddress2"`
	BillingAddress3  string `xorm:"not null VARCHAR(200)" json:"billingAddress3,omitempty" xml:"billingAddress3"`
	BillingAddress4  string `xorm:"not null VARCHAR(200)" json:"billingAddress4,omitempty" xml:"billingAddress4"`
	BillingState     string `xorm:"not null VARCHAR(100)" json:"billingState,omitempty" xml:"billingState"`
	BillingPostcode  string `xorm:"not null VARCHAR(100)" json:"billingPostcode,omitempty" xml:"billingPostcode"`
	BillingCountry   string `xorm:"not null VARCHAR(100)" json:"billingCountry,omitempty" xml:"billingCountry"`
	ShippingAddress1 string `xorm:"not null VARCHAR(200)" json:"shippingAddress1,omitempty" xml:"shippingAddress1"`
	ShippingAddress2 string `xorm:"not null VARCHAR(200)" json:"shippingAddress2,omitempty" xml:"shippingAddress2"`
	ShippingAddress3 string `xorm:"not null VARCHAR(200)" json:"shippingAddress3,omitempty" xml:"shippingAddress3"`
	ShippingAddress4 string `xorm:"not null VARCHAR(200)" json:"shippingAddress4,omitempty" xml:"shippingAddress4"`
	ShippingState    string `xorm:"not null VARCHAR(100)" json:"shippingState,omitempty" xml:"shippingState"`
	ShippingPostcode string `xorm:"not null VARCHAR(100)" json:"shippingPostcode,omitempty" xml:"shippingPostcode"`
	ShippingCountry  string `xorm:"not null VARCHAR(100)" json:"shippingCountry,omitempty" xml:"shippingCountry"`
	BranchArea       string `xorm:"not null VARCHAR(500)" json:"branchArea,omitempty" xml:"branchArea"`
	BranchRemark     string `xorm:"not null VARCHAR(500)" json:"branchRemark,omitempty" xml:"branchRemark"`
	BranchActive     int    `xorm:"not null default 1 comment('1=active, 0=not active') INT" json:"branchActive,omitempty" xml:"branchActive"`
	BranchUdf        string `xorm:"JSON" json:"branchUdf,omitempty" xml:"branchUdf"`
	Company          string `xorm:"default '' VARCHAR(50)" json:"company,omitempty" xml:"company"`
}
