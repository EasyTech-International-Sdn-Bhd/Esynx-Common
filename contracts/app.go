package contracts

import (
	"context"
	"github.com/hibiken/asynq"
)

type IAppOneWayService interface {
	ActOnEvents() error
	RetrieveUptoFuture() ([]*asynq.TaskInfo, error)
	RetrieveOnlyUpdated() ([]*asynq.TaskInfo, error)
	RetrieveSpecific(data interface{}) ([]*asynq.TaskInfo, error)
	GetHandlers() map[string]func(context.Context, *asynq.Task) error
}

type IAppBidirectionalService interface {
	ActOnEvents() error
	RetrieveUptoFuture() ([]*asynq.TaskInfo, error)
	RetrieveOnlyUpdated() ([]*asynq.TaskInfo, error)
	RetrieveSpecific(data interface{}) ([]*asynq.TaskInfo, error)
	GetHandlers() map[string]func(context.Context, *asynq.Task) error
	Sync() ([]*asynq.TaskInfo, error)
}
