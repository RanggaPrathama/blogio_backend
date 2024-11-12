package interfaces

import (
	"blogio/internal/domain/entity"
	"context"
)

type UserInterface interface {
	FindAll(context.Context) ([]entity.User, error)
	FindByID(context.Context, string) (entity.User, error)
}