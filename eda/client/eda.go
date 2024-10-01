package client

import (
	"context"
	"github.com/easytech-international-sdn-bhd/esynx-common/eda/events"
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
	Produce(HandlerType, request.EdaHeader, map[string]string, string) error
	GetConsumer(string, HandlerType) (IEventDrivenMessageConsumer, error)
	Close() error
}

type IEventDrivenMessageConsumer interface {
	Create(subscribers []SubscriberConfig) (map[events.EDARoutes]IEventDrivenMessageHandler, error)
}

type IEventDrivenMessageHandler interface {
	Consume(c context.Context, ch chan<- *MessageEDA) error
	UnderlyingEngine() interface{}
}
