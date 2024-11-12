package service

import (
	"blogio/internal/domain/entity"
	"blogio/internal/domain/repository/interfaces"
	"blogio/internal/service/responses"
	"context"
)

type UserService interface {
	FindAll(ctx context.Context) ([]entity.User, error)
	FindByID(ctx context.Context,id string) (entity.User, error)
}

type User_service struct {
	userRepo interfaces.UserRepository
}

func NewUserService(userRepo interfaces.UserRepository) UserService {
	return &User_service{
		userRepo: userRepo,
	}
}


func (u *User_service) FindAll(c context.Context) ([]entity.User, error) {
	users, err := u.userRepo.FindAll(c)
	if err != nil {
		return users, responses.NewErrorNotFound(err)
		}

	return users, err
}


func (u *User_service) FindByID(c context.Context, id string) (entity.User, error) {
	users, err := u.userRepo.FindByID(c, id)
	if err != nil {
		return users, responses.NewErrorNotFound(err)
	}

	return users, err
}