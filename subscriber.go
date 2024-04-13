package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

type MessageConsumer struct {
	redisClient  Redis
	subscription *redis.PubSub
}

func NewMessageConsumer(redis Redis) *MessageConsumer {
	return &MessageConsumer{
		redisClient: redis,
	}
}

func (c *MessageConsumer) ConsumerMessages(ctx context.Context, queueNames []string) {
	for _, queueName := range queueNames {
		go c.handleCustomType1Logic(ctx, queueName)
	}
}

func (c *MessageConsumer) handleCustomType1Logic(ctx context.Context, queueName string) {
	consumerCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Printf("[%s] Consumer started listening...\n", queueName)

	c.subscription = c.redisClient.RedisClient.Subscribe(queueName)
	defer c.subscription.Close()

	channel := c.subscription.Channel()

	for {
		select {
		case <-consumerCtx.Done():
			log.Printf("[%s] Consumer stopped listening...\n", queueName)
			return
		case msg := <-channel:
			var messageObj interface{}
			err := json.Unmarshal([]byte(msg.Payload), &messageObj)
			if err != nil {
				log.Printf("[%s] Failed to deserialize message: %v", queueName, err)
				continue
			}

			fmt.Printf("[%s] Received message: %+v\n", queueName, messageObj)
		}
	}
}
