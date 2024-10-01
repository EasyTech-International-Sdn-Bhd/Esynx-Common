package client

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-common/eda/events"
	"strings"
	"time"
)

type SubscriberConfig struct {
	Name   events.EDARoutes
	Filter string
}

type GcpConsumer struct {
	ctx         context.Context
	handlerType HandlerType
	client      *pubsub.Client
	clientId    string
	projectId   string
	topic       *pubsub.Topic
}

func NewGcpConsumer(ctx context.Context, clientId, projectId string, handlerType HandlerType, client *pubsub.Client, topic *pubsub.Topic) *GcpConsumer {
	return &GcpConsumer{
		ctx:         ctx,
		client:      client,
		clientId:    clientId,
		projectId:   projectId,
		topic:       topic,
		handlerType: handlerType,
	}
}

func (c *GcpConsumer) Create(subscriberConfig []SubscriberConfig) (map[events.EDARoutes]IEventDrivenMessageHandler, error) {
	handlers := make(map[events.EDARoutes]IEventDrivenMessageHandler)
	for _, subscriber := range subscriberConfig {
		subId := strings.ToLower(fmt.Sprintf("sub.%s.%s.%s", c.clientId, subscriber.Name, c.handlerType.String()))
		subscription := c.client.SubscriptionInProject(subId, c.projectId)
		exists, err := subscription.Exists(c.ctx)
		if err != nil {
			return nil, err
		}
		if !exists {
			subscription, err = c.client.CreateSubscription(c.ctx, subId, pubsub.SubscriptionConfig{
				Topic:                 c.topic,
				AckDeadline:           10 * time.Minute,
				EnableMessageOrdering: true,
				Filter:                subscriber.Filter,
			})
			if err != nil {
				return nil, err
			}
		}

		handlers[subscriber.Name] = NewGcpMessageHandler(subscription)
	}

	return handlers, nil
}
