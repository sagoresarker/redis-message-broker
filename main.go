package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	redisClient := NewRedis()
	defer redisClient.RedisClient.Close()

	publisher := NewMessagePublisher(redisClient)

	subscriber := NewMessageConsumer(redisClient)

	go subscriber.ConsumerMessages(ctx, []string{"test"})

	publisher.PublishMessages(ctx, "Hello, Redis!", "test")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
