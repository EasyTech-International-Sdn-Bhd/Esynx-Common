package client

import (
	"context"
	"github.com/easytech-international-sdn-bhd/esynx-common/eda/request"
)

type HandlerType string

func (h HandlerType) String() string {
	return string(h)
}

const (
	ServerToAccounting HandlerType = "server.acc"
	ClientToServer                 = "client.server"
)

type IEventDrivenMessageProducer interface {
	Produce(HandlerType, request.EdaHeader, interface{}, map[string]string, string) error
	GetConsumer(string, HandlerType) (IEventDrivenMessageConsumer, error)
	Close() error
}

type IEventDrivenMessageConsumer interface {
	Create(handlerType HandlerType, subscribers []SubscriberConfig) (map[string]IEventDrivenMessageHandler, error)
}

type IEventDrivenMessageHandler interface {
	Consume(context.Context, func(context.Context, *MessageEDA)) error
}
