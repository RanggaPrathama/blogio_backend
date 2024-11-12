package service

import (
	"blogio/internal/domain/entity"
	Uservice "blogio/internal/service/interfaces"
	Urepository "blogio/internal/domain/repository/interfaces"
	"blogio/internal/service/responses"
	"context"
)



type User_service struct {
	userRepo Urepository.UserRepository 
}

func NewUserService(userRepo Urepository.UserRepository)  Uservice.UserInterface {
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


func (u *User_service) CreateUser(c context.Context) (entity.User, error) {
	users, err := u.userRepo.CreateUser(c)
	if err != nil {
		return users, responses.NewErrorNotFound(err)
	}

	return users, err
}