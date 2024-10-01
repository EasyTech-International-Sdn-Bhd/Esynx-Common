package request

type EdaDataType string

const (
	EDADataTypeAgent             EdaDataType = "Agent"
	EDADataTypeCreditNote        EdaDataType = "CreditNote"
	EDADataTypeCreditNoteDetails EdaDataType = "CreditNoteDetails"
	EDADataTypeInvoice           EdaDataType = "Invoice"
	EDADataTypeInvoiceDetails    EdaDataType = "InvoiceDetails"
	EDADataTypeDebitNote         EdaDataType = "DebitNote"
	EDADataTypeDebitNoteDetails  EdaDataType = "DebitNoteDetails"
	EDADataTypeDebtor            EdaDataType = "Debtor"
	EDADataTypeDebtorBranch      EdaDataType = "DebtorBranch"
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
	DataType   EdaDataType   `json:"source"`
	TaskID     string        `json:"task_id"`
	Username   string        `json:"username"`
	ClientID   string        `json:"client_id"`
	AppName    string        `json:"app_name"`
	ActionType EdaActionType `json:"action_type"`
	ActionTime int64         `json:"action_time"`
	Metadata   interface{}   `json:"metadata"`
	Data       interface{}   `json:"data"`
}
