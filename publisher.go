package main

import (
	"context"
	"encoding/json"
	"log"
)

type MessagePublisher struct {
	redisClient Redis
}

func NewMessagePublisher(redisClient Redis) *MessagePublisher {
	return &MessagePublisher{redisClient}
}

func (p *MessagePublisher) PublishMessages(ctx context.Context, message interface{}, queueName string) {
	serializedMessage, err := json.Marshal(message)
	if err != nil {
		log.Printf("[%s] Failed to serialize message: %v", queueName, err)
		return
	}

	err = p.redisClient.RedisClient.Publish(queueName, serializedMessage).Err()
	if err != nil {
		log.Printf("[%s] Failed to publish message: %v", queueName, err)
	}
}
