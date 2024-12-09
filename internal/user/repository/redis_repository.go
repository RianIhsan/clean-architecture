package repository

import (
	"context"
	"encoding/json"
	"github.com/RianIhsan/go-clean-architecture-v2/internal/user"
	"github.com/RianIhsan/go-clean-architecture-v2/internal/user/model"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type userRedisRepository struct {
	redisClient *redis.Client
}

func NewUserRedisRepository(redisClient *redis.Client) user.UserRedisRepository {
	return &userRedisRepository{
		redisClient: redisClient,
	}
}

// Set cache user
func (u *userRedisRepository) Set(ctx context.Context, key string, expiration time.Duration, value *model.User) error {
	userBytes, err := json.Marshal(value)
	if err != nil {
		return errors.Wrap(err, "UserRedisRepository.Set.json.Marshal")
	}

	if err := u.redisClient.Set(ctx, key, userBytes, time.Second*expiration).Err(); err != nil {
		return errors.Wrap(err, "UserRedisRepository.Set.redisClient.Set")
	}
	return nil
}

// Get cache data
func (u *userRedisRepository) Get(ctx context.Context, key string) (*model.User, error) {
	userBytes, err := u.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "UserRedisRepository.Get.redisClient.Get")
	}

	user := new(model.User)
	if err := json.Unmarshal(userBytes, user); err != nil {
		return nil, errors.Wrap(err, "UserRedisRepository.Get.Unmarshall")
	}
	return user, nil
}

// Delete cache data
func (u *userRedisRepository) Delete(ctx context.Context, key string) error {
	if err := u.redisClient.Del(ctx, key).Err(); err != nil {
		return errors.Wrap(err, "UserRedisRepository.Delete.redisClient.Del")
	}
	return nil
}
