package request

type DataType string

const (
	EDADataTypeAgent             DataType = "Agent"
	EDADataTypeCreditNote        DataType = "CreditNote"
	EDADataTypeCreditNoteDetails DataType = "CreditNoteDetails"
	EDADataTypeInvoice           DataType = "Invoice"
	EDADataTypeInvoiceDetails    DataType = "InvoiceDetails"
	EDADataTypeDebitNote         DataType = "DebitNote"
	EDADataTypeDebitNoteDetails  DataType = "DebitNoteDetails"
	EDADataTypeCashSales         DataType = "CashSales"
	EDADataTypeCashSalesDetails  DataType = "CashSalesDetails"
	EDADataTypeDebtor            DataType = "Debtor"
	EDADataTypeDebtorBranch      DataType = "DebtorBranch"
)

type EdaHeader struct {
	DataType   DataType    `json:"source"`
	TaskID     string      `json:"task_id"`
	Username   string      `json:"username"`
	ClientID   string      `json:"client_id"`
	AppName    string      `json:"app_name"`
	ActionType string      `json:"action_type"`
	ActionTime string      `json:"action_time"`
	Metadata   interface{} `json:"metadata"`
	Data       interface{} `json:"data"`
}
