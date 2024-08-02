package contracts

import "context"

type AppOpsContext struct {
	MsgChan  chan interface{}
	Context  context.Context
	Shutdown chan struct{}
	Cancel   context.CancelFunc
}

type IAppOneWayService interface {
	RetrieveUptoFuture(ctx AppOpsContext) error
	ActOnEvents(ctx AppOpsContext) error
	RetrieveOnlyUpdated(ctx AppOpsContext) error
	RetrieveSpecific(ctx AppOpsContext, data interface{}) error
	GetLogs(ctx AppOpsContext) []interface{}
}

type IAppBidirectionalService interface {
	RetrieveUptoFuture(ctx AppOpsContext) error
	ActOnEvents(ctx AppOpsContext) error
	RetrieveOnlyUpdated(ctx AppOpsContext) error
	RetrieveSpecific(ctx AppOpsContext, data interface{}) error
	GetLogs(ctx AppOpsContext) []interface{}
	Sync(ctx AppOpsContext) error
}
