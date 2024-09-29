package client

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-common/eda/request"
	"github.com/goccy/go-json"
	"google.golang.org/api/option"
	"strings"
	"sync"
	"time"
)

type topicPair struct {
	Handler  HandlerType
	ClientId string
	TopicId  string
	Topic    *pubsub.Topic
}

type GcpBroker struct {
	ctx            context.Context
	client         *pubsub.Client
	topics         []topicPair
	projectId      string
	projectOptions option.ClientOption
	mu             sync.Mutex
}

func NewGcpBroker(gcpClientOption option.ClientOption, gcpProjectID string) (*GcpBroker, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, gcpProjectID, gcpClientOption)
	if err != nil {
		return nil, err
	}
	return &GcpBroker{
		ctx:            ctx,
		client:         client,
		topics:         make([]topicPair, 0),
		projectId:      gcpProjectID,
		projectOptions: gcpClientOption,
	}, nil
}

func (b *GcpBroker) Produce(ht HandlerType, req request.EdaHeader, data interface{}, attr map[string]string, orderBy string) error {
	topicId := strings.ToLower(fmt.Sprintf("%s.%s", ht.String(), req.ClientID))
	topic, err := b.getOrCreateTopic(req.ClientID, topicId, ht)
	if err != nil {
		return err
	}
	req.Data = data
	dt, err := json.Marshal(req)
	if err != nil {
		return err
	}
	result := topic.Publish(b.ctx, &pubsub.Message{
		Data:        dt,
		Attributes:  attr,
		OrderingKey: orderBy,
	})
	_, err = result.Get(b.ctx)
	if err != nil {
		return fmt.Errorf("PublishEvent.Get: %w", err)
	}
	return nil
}

func (b *GcpBroker) GetConsumer(clientId string, ht HandlerType) (IEventDrivenMessageConsumer, error) {
	topicId := strings.ToLower(fmt.Sprintf("%s.%s", ht.String(), clientId))
	var topic *pubsub.Topic
	b.mu.Lock()
	for _, t := range b.topics {
		if t.ClientId == clientId && t.TopicId == topicId {
			topic = t.Topic
			break
		}
	}
	b.mu.Unlock()
	if topic == nil {
		return nil, fmt.Errorf("topic not found")
	}
	return NewGcpConsumer(b.ctx, clientId, b.projectId, b.client, topic), nil
}

func (b *GcpBroker) getOrCreateTopic(clientId, topicId string, ht HandlerType) (*pubsub.Topic, error) {
	b.mu.Lock()
	for _, t := range b.topics {
		if t.TopicId == topicId {
			removeTopic := func() {
				for i, tp := range b.topics {
					if tp.TopicId == topicId {
						b.topics = append(b.topics[:i], b.topics[i+1:]...)
						break
					}
				}
			}
			exists, err := t.Topic.Exists(b.ctx)
			if err != nil {
				return nil, err
			}
			if !exists {
				removeTopic()
				break
			}
			b.mu.Unlock()
			return t.Topic, nil
		}
	}
	b.mu.Unlock()

	var topic = b.client.TopicInProject(topicId, b.projectId)
	exists, err := topic.Exists(b.ctx)
	if err != nil {
		return nil, err
	}
	if !exists {
		topic, err = b.client.CreateTopic(b.ctx, topicId)
		if err != nil {
			return nil, err
		}
	}
	topic.EnableMessageOrdering = true

	subId := fmt.Sprintf("sub.def.%s", topicId)
	subscriber := b.client.SubscriptionInProject(subId, b.projectId)
	exists, err = subscriber.Exists(b.ctx)
	if err != nil {
		return nil, err
	}
	if !exists {
		_, err = b.client.CreateSubscription(b.ctx, subId, pubsub.SubscriptionConfig{
			Topic:                 topic,
			AckDeadline:           600 * time.Second,
			EnableMessageOrdering: true,
		})
		if err != nil {
			return nil, fmt.Errorf("CreateSubscriptionWithConfig: %w", err)
		}
	}

	b.mu.Lock()
	b.topics = append(b.topics, topicPair{
		Handler:  ht,
		TopicId:  topicId,
		Topic:    topic,
		ClientId: clientId,
	})
	b.mu.Unlock()

	return topic, nil
}

func (b *GcpBroker) Close() error {
	return b.client.Close()
}
