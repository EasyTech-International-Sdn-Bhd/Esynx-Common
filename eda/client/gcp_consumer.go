package client

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

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

func (c *GcpConsumer) Create(handlerType HandlerType, subscribers []string, filter string) (map[string]IEventDrivenMessageHandler, error) {
	handlers := make(map[string]IEventDrivenMessageHandler)
	errChan := make(chan error, len(subscribers))
	var mu sync.Mutex

	var wg sync.WaitGroup
	for _, subscriber := range subscribers {
		wg.Add(1)
		go func(subscriber string) {
			defer wg.Done()

			subId := strings.ToLower(fmt.Sprintf("sub.%s.%s", subscriber, handlerType.String()))
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
					Filter:                filter,
				})
				if err != nil {
					errChan <- err
					return
				}
			}

			mu.Lock()
			handlers[subscriber] = NewGcpMessageHandler(subscription)
			mu.Unlock()

			errChan <- nil
		}(subscriber)
	}

	wg.Wait()

	close(errChan)

	for i := 0; i < len(subscribers); i++ {
		if err := <-errChan; err != nil {
			return nil, err
		}
	}

	return handlers, nil
}
