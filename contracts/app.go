package contracts

import (
	"context"
	"github.com/hibiken/asynq"
)

type IAppOneWayService interface {
	RetrieveUptoFuture() error
	ActOnEvents() error
	RetrieveOnlyUpdated() error
	RetrieveSpecific(data interface{}) error
	GetHandlers() map[string]func(context.Context, *asynq.Task) error
}

type IAppBidirectionalService interface {
	RetrieveUptoFuture() error
	ActOnEvents() error
	RetrieveOnlyUpdated() error
	RetrieveSpecific(data interface{}) error
	GetHandlers() map[string]func(context.Context, *asynq.Task) error
	Sync() error
}
