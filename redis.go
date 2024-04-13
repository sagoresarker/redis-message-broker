package main

import "github.com/go-redis/redis"

type Redis struct {
	RedisClient *redis.Client
}

func NewRedis() Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return Redis{
		RedisClient: client,
	}
}
