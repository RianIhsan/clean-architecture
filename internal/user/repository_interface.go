package user

import (
	"context"
	"github.com/RianIhsan/go-clean-architecture-v2/internal/user/model"
	"time"
)

// UserPostgresRepository defines methods the services layer expects.
// any repositories it interacts with to implement
type UserPostgresRepository interface {
	Create(ctx context.Context, entity *model.User) (*model.User, error)

	Update(ctx context.Context, entity *model.User) (*model.User, error)

	FindByEmail(ctx context.Context, entity *model.User) (*model.User, error)

	FindById(ctx context.Context, entity *model.User) (*model.User, error)

	FindAlreadyExistByEmail(ctx context.Context, entity *model.User) (int64, error)
}

// UserRedisRepository defines methods the services layer expects.
// any repositories it interacts with to implement.
type UserRedisRepository interface {
	Set(ctx context.Context, key string, expiration time.Duration, value *model.User) error

	Get(ctx context.Context, key string) (*model.User, error)

	Delete(ctx context.Context, key string) error
}
