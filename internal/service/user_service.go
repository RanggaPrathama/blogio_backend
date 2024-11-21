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


func (u *User_service) UpdateUser(c context.Context, id string, user entity.User) (entity.User, error){

	hashedPassword, err := helper.GeneratePassword(user.PASSWORD)
	if err != nil {
		return user, responses.NewErrorNotFound(err)
	}

	user.PASSWORD = hashedPassword

	usersID, err := u.FindByID(c, id)

	if err != nil {
		return usersID, responses.NewErrorNotFound(fmt.Errorf("user not found"))
	}

	id = usersID.ID.Hex()

	users, err := u.userRepo.UpdateUser(c, id, user)

	if err != nil {
		return users, responses.NewErrorNotFound(err)
	}

	return users, err 

}

func (u *User_service) DeleteUser(c context.Context, id string) error {

	usersID, err := u.userRepo.FindByID(c, id)

	fmt.Println("USERS ID", usersID)

	if err != nil {
		return responses.NewErrorNotFound(fmt.Errorf("user not found"))
	}

	err = u.userRepo.DeleteUser(c, usersID.ID.Hex())

	fmt.Println("ERROR", err)
	if err != nil {
		return responses.NewErrorNotFound(err)
	}

	return err

}