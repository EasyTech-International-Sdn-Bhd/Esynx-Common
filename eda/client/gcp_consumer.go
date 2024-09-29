package client

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-common/eda/events"
	"strings"
	"sync"
	"time"
)

type SubscriberConfig struct {
	Name   events.EDARoutes
	Filter string
}

type GcpConsumer struct {
	ctx       context.Context
	client    *pubsub.Client
	clientId  string
	projectId string
	topic     *pubsub.Topic
}

func NewGcpConsumer(ctx context.Context, clientId, projectId string, client *pubsub.Client, topic *pubsub.Topic) *GcpConsumer {
	return &GcpConsumer{
		ctx:       ctx,
		client:    client,
		clientId:  clientId,
		projectId: projectId,
		topic:     topic,
	}
}

func (c *GcpConsumer) Create(handlerType HandlerType, subscriberConfig []SubscriberConfig) (map[events.EDARoutes]IEventDrivenMessageHandler, error) {
	handlers := make(map[events.EDARoutes]IEventDrivenMessageHandler)
	errChan := make(chan error, len(subscriberConfig))
	var mu sync.Mutex

	var wg sync.WaitGroup
	for _, subscriber := range subscriberConfig {
		wg.Add(1)
		go func(subscriber SubscriberConfig) {
			defer wg.Done()

			subId := strings.ToLower(fmt.Sprintf("sub.%s.%s", subscriber.Name, handlerType.String()))
			subscription := c.client.SubscriptionInProject(subId, c.projectId)
			exists, err := subscription.Exists(c.ctx)
			if err != nil {
				errChan <- err
				return
			}
			if !exists {
				subscription, err = c.client.CreateSubscription(c.ctx, subId, pubsub.SubscriptionConfig{
					Topic:                 c.topic,
					AckDeadline:           10 * time.Minute,
					EnableMessageOrdering: true,
					Filter:                subscriber.Filter,
				})
				if err != nil {
					errChan <- err
					return
				}
			}

			mu.Lock()
			handlers[subscriber.Name] = NewGcpMessageHandler(subscription)
			mu.Unlock()

			errChan <- nil
		}(subscriber)
	}

	wg.Wait()

	close(errChan)

	for i := 0; i < len(subscriberConfig); i++ {
		if err := <-errChan; err != nil {
			return nil, err
		}
	}

	return handlers, nil
}
