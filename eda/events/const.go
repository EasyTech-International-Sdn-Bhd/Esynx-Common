package events

type EDARoutes string

const (
	EDARoute_Agent        EDARoutes = "Agent"
	EDARoute_Invoice      EDARoutes = "Invoice"
	EDARoute_CreditNote   EDARoutes = "CreditNote"
	EDARoute_DebitNote    EDARoutes = "DebitNote"
	EDARoute_Debtor       EDARoutes = "Debtor"
	EDARoute_CashSales    EDARoutes = "CashSales"
	EDARoute_ArInvoice    EDARoutes = "ArInvoice"
	EDARoute_ArCreditNote EDARoutes = "ArCreditNote"
	EDARoute_ArDebitNote  EDARoutes = "ArDebitNote"
)

type EDADataSource string

const (
	EDADataSource_Invoice           EDADataSource = "Invoice"
	EDADataSource_InvoiceDetails    EDADataSource = "InvoiceDetails"
	EDADataSource_DebitNote         EDADataSource = "DebitNote"
	EDADataSource_DebitNoteDetails  EDADataSource = "DebitNoteDetails"
	EDADataSource_CreditNote        EDADataSource = "CreditNote"
	EDADataSource_CreditNoteDetails EDADataSource = "CreditNoteDetails"
	EDADataSource_Debtor            EDADataSource = "Debtor"
	EDADataSource_DebtorBranch      EDADataSource = "DebtorBranch"
	EDADataSource_Agent             EDADataSource = "Agent"
)
