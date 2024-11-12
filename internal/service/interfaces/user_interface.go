package interfaces

import (
	"blogio/internal/domain/entity"
	"context"
)

type UserInterface interface {
	FindAll(ctx context.Context) ([]entity.User, error)
	FindByID(ctx context.Context, id string) (entity.User, error)
	CreateUser(ctx context.Context) (entity.User, error)
}