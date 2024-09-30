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

type EdaActionType string

const (
	// EDAActionTypeSave represents the action type for saving an entity from accounting software.
	EDAActionTypeSave   EdaActionType = "Save"
	EDAActionTypeCreate EdaActionType = "Create"
	EDAActionTypeUpdate EdaActionType = "Update"
	EDAActionTypeDelete EdaActionType = "Delete"
)

type EdaHeader struct {
	DataType   DataType      `json:"source"`
	TaskID     string        `json:"task_id"`
	Username   string        `json:"username"`
	ClientID   string        `json:"client_id"`
	AppName    string        `json:"app_name"`
	ActionType EdaActionType `json:"action_type"`
	ActionTime int64         `json:"action_time"`
	Metadata   interface{}   `json:"metadata"`
	Data       interface{}   `json:"data"`
}
