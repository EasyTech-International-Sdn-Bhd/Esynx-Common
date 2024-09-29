package forms

type DataType string

const (
	EDADataTypeAgent        DataType = "Agent"
	EDADataTypeCreditNote   DataType = "CreditNote"
	EDADataTypeInvoice      DataType = "Invoice"
	EDADataTypeDebitNote    DataType = "DebitNote"
	EDADataTypeDebtor       DataType = "Debtor"
	EDADataTypeDebtorBranch DataType = "DebtorBranch"
)

type DefaultForm struct {
	DataType   DataType    `json:"source"`
	TaskID     string      `json:"task_id"`
	Username   string      `json:"username"`
	ClientID   string      `json:"client_id"`
	AppName    string      `json:"app_name"`
	ActionType string      `json:"action_type"`
	ActionTime string      `json:"action_time"`
	Metadata   interface{} `json:"metadata"`
}
