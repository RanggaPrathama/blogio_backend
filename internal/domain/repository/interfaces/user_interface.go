package interfaces

import (
	"blogio/internal/domain/entity"
	"context"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]entity.User, error)
	FindByID(ctx context.Context, id string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	FindByEmail(ctx context.Context, email string) (entity.User, error)
	UpdateUser(ctx context.Context, id string, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id string) error
}
