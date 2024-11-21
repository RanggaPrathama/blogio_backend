package service

import (
	"blogio/config"
	"blogio/helper"
	"blogio/internal/domain/entity"
	AService"blogio/internal/service/interfaces"
	Urepository "blogio/internal/domain/repository/interfaces"
	"blogio/internal/service/responses"
	"context"
	"fmt"
)


type AuthServices struct {
	UserRepository Urepository.UserRepository
}

func NewAuthService(userRepo Urepository.UserRepository) AService.AuthInterface{
	return &AuthServices{
		UserRepository: userRepo,
	}
}

func (auth *AuthServices) Login( ctx context.Context, email string, password string) (entity.LoginStruct, error){
	
	var loginStruct entity.LoginStruct

	user, err := auth.UserRepository.FindByEmail(ctx, email)
	
	if err != nil {
		fmt.Println("error", err)
		return loginStruct, responses.NewErrorNotFound(err)
	}

	if user.EMAIL == "" {
		return loginStruct, responses.NewErrorNotFound(fmt.Errorf("user not found"))
	}

	err = helper.ComparePassword(user.PASSWORD, password)

	if err != nil {
		return loginStruct, responses.NewErrorNotFound(fmt.Errorf("password not match"))
	}

	secretKey := config.LoadEnv("JWT_SECRET_KEY")

	token, err := helper.GenerateToken(user.ID.Hex(), user.EMAIL, secretKey)

	if err != nil || token == "" {
		return loginStruct, responses.NewErrorNotFound(fmt.Errorf("error generating token"))
	}

	loginStruct.Token = token
	loginStruct.User = user

	return loginStruct, err


}