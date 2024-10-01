package client

import (
	"cloud.google.com/go/pubsub"
	"context"
	"time"
)

type MessageEDA struct {
	Data       []byte
	Attributes map[string]string
	Ack        func()
	Nack       func()
}

type GcpMessageHandler struct {
	sub *pubsub.Subscription
}

func NewGcpMessageHandler(sub *pubsub.Subscription) *GcpMessageHandler {
	return &GcpMessageHandler{sub: sub}
}

func (m *GcpMessageHandler) Consume(c context.Context, ch chan<- *MessageEDA) error {
	m.sub.ReceiveSettings.MaxOutstandingMessages = -1
	m.sub.ReceiveSettings.MaxOutstandingBytes = -1
	m.sub.ReceiveSettings.MaxExtensionPeriod = time.Second * 10
	return m.sub.Receive(c, func(ctx context.Context, msg *pubsub.Message) {
		newMsg := &MessageEDA{
			Data:       msg.Data,
			Attributes: msg.Attributes,
			Ack: func() {
				msg.Ack()
			},
			Nack: func() {
				msg.Nack()
			},
		}
		ch <- newMsg
	})
}

func (m *GcpMessageHandler) UnderlyingEngine() interface{} {
	return m.sub
}
