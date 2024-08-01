package contracts

type IAppOneWayService interface {
	RetrieveUptoFuture() error
	ActOnEvents() error
	RetrieveOnlyUpdated() error
	RetrieveSpecific(data interface{}) error
	GetLogs() []interface{}
}

type IAppBidirectionalService interface {
	RetrieveUptoFuture() error
	ActOnEvents() error
	RetrieveOnlyUpdated() error
	RetrieveSpecific(data interface{}) error
	GetLogs() []interface{}
	Sync() error
}
