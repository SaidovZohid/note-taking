package storage

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"
)

type InMemoryStorageI interface {
	Set(key, val string, exp time.Duration) error 
	Get(key string) (string, error)
}

type storageRedis struct {
	client *redis.Client
}

func NewRedisStorage(rdb *redis.Client) InMemoryStorageI {
	return &storageRedis{
		client: rdb,
	}
}

func (rd *storageRedis) Set(key, val string, exp time.Duration) error {
	err := rd.client.Set(context.Background(), key, val, exp).Err()
	if err != nil {
		return err
	}

	return nil
}

func (rd *storageRedis) Get(key string) (string, error) {
	val, err := rd.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}