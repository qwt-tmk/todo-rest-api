package kvs

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCommandar struct {
	cli *redis.Client
}

func NewRedisCommandar() *RedisCommandar {
	return &RedisCommandar{
		cli: GetRedisClient(),
	}
}

func (rc *RedisCommandar) Save(ctx context.Context, duration time.Duration, userID, jwtID string) error {
	// Redisクライアントがnilの場合はエラーを返す
	if rc.cli == nil {
		return errors.New("failed to get Redis client")
	}
	status := rc.cli.Set(ctx, userID, jwtID, duration)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

func (rc *RedisCommandar) Load(ctx context.Context, userID string) (string, error) {
	// Redisクライアントがnilの場合はエラーを返す
	if rc.cli == nil {
		return "", errors.New("failed to get Redis client")
	}
	status := rc.cli.Get(ctx, userID)
	if status.Err() != nil {
		// nilだったら空文字を返す
		if status.Err() == redis.Nil {
			return "", nil
		}
		return "", status.Err()
	}
	return status.Val(), nil
}

func (rc *RedisCommandar) Delete(ctx context.Context, userID string) error {
	// Redisクライアントがnilの場合はエラーを返す
	if rc.cli == nil {
		return errors.New("failed to get Redis client")
	}
	status := rc.cli.Del(ctx, userID)
	if status.Err() != nil {
		if status.Err() == redis.Nil {
			return nil
		}
		return status.Err()
	}
	return nil
}
