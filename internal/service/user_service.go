package service

import (
	"blogio/helper"
	"blogio/internal/domain/entity"
	Urepository "blogio/internal/domain/repository/interfaces"
	Uservice "blogio/internal/service/interfaces"
	"blogio/internal/service/responses"
	"context"
	"fmt"
	// "go.mongodb.org/mongo-driver/mongo"
)

type User_service struct {
	userRepo Urepository.UserRepository
}

func NewUserService(userRepo Urepository.UserRepository) Uservice.UserInterface {
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

func (u *User_service) CreateUser(c context.Context, user entity.User) (entity.User, error) {

	existUser, _ := u.userRepo.FindByEmail(c, user.EMAIL)
	 fmt.Println(existUser)
	// if err != nil && err != mongo.ErrNoDocuments {
	// 	return user, responses.NewErrorNotFound(err)
	// }

	if existUser.EMAIL == user.EMAIL {
		return user,  responses.NewErrorNotFound(fmt.Errorf("email already exists"))
	}

	hashedPassword, err := helper.GeneratePassword(user.PASSWORD)
	if err != nil {
		return user, responses.NewErrorNotFound(err)
	}

	user.PASSWORD = hashedPassword

	users, err := u.userRepo.CreateUser(c, user)
	if err != nil {
		return users, responses.NewErrorNotFound(err)
	}

	return users, err
}
