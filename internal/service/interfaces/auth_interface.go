package interfaces

import (
	"blogio/internal/domain/entity"
	"context"
)

type AuthInterface interface {
	Login(ctx context.Context, email string, password string) (entity.LoginStruct, error)
}