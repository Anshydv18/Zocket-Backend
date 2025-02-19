package services

import (
	"backend/models/dto"
	"backend/models/entity"
	"backend/models/requests"
	"backend/utils"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func UserLogin(ctx *context.Context, request *requests.UserLoginRequest) (*dto.UserDto, error) {
	if !utils.ValidateEmail(request.Email) {
		return nil, errors.New("not a valid email")
	}

	data, err := entity.GetUserInfoByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	isUserPasswordMatched := checkPasswordHash(request.Password, data.Password)
	if !isUserPasswordMatched {
		return nil, errors.New("Email and Password don't match")
	}

	return &dto.UserDto{
		Email:        data.Email,
		Name:         data.Name,
		TaskAssigned: data.TaskAssigned,
	}, nil
}

func CreateUserProfile(ctx *context.Context, request *requests.UserProfileRequest) error {
	// if !utils.ValidateEmail(request.Email) {
	// 	return errors.New("not a valid email")
	// }

	request.Password, _ = HassPassword(request.Password)

	err := entity.CreateUserProfile(ctx, request)
	return err
}

func HassPassword(password string) (string, error) {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hasedPasswordStr := string(hasedPassword)
	if err != nil {
		return hasedPasswordStr, err
	}

	return hasedPasswordStr, nil
}

func checkPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
