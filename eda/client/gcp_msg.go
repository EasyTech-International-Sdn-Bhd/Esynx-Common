package client

import (
	"cloud.google.com/go/pubsub"
	"context"
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

func (m *GcpMessageHandler) Consume(ctx context.Context, cb func(context.Context, *MessageEDA)) error {
	return m.sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		newMsg := &MessageEDA{
			Data:       msg.Data,
			Attributes: msg.Attributes,
			Ack:        msg.Ack,
			Nack:       msg.Nack,
		}
		cb(ctx, newMsg)
	})
}
