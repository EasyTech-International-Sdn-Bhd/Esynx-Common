package events

type EDARoutes string

const (
	EDARoute_Invoice           EDARoutes = "Invoice"
	EDARoute_InvoiceDetails    EDARoutes = "InvoiceDetails"
	EDARoute_CreditNote        EDARoutes = "CreditNote"
	EDARoute_CreditNoteDetails EDARoutes = "CreditNoteDetails"
	EDARoute_DebitNote         EDARoutes = "DebitNote"
	EDARoute_DebitNoteDetails  EDARoutes = "DebitNoteDetails"
	EDARoute_Debtor            EDARoutes = "Debtor"
	EDARoute_DebtorBranch      EDARoutes = "DebtorBranch"
	EDARoute_Agent             EDARoutes = "Agent"
)
