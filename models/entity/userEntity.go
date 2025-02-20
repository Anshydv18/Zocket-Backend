package entity

import (
	"backend/base"
	"backend/constants"
	"backend/models/requests"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

type UserData struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"Password"`
	TaskAssigned interface{}
}

func GetUserInfoByEmail(ctx *context.Context, email string) (*UserData, error) {
	dbClient := base.DBInstance
	if dbClient == nil {
		return nil, errors.New("failed while connecting to database")
	}

	collection := dbClient.Database(constants.ZOCKETDB).Collection(constants.USER_COLLECTION)

	filter := bson.M{
		"email": email,
	}

	var UserData UserData
	data := collection.FindOne(*ctx, filter)

	if err := data.Decode(&UserData); err != nil {
		return &UserData, err
	}

	return &UserData, nil
}

func CreateUserProfile(ctx *context.Context, request *requests.UserProfileRequest) error {
	dbClient := base.DBInstance
	if dbClient == nil {
		return errors.New("failed while connecting to database")
	}

	collection := dbClient.Database(constants.ZOCKETDB).Collection(constants.USER_COLLECTION)

	userData := bson.D{
		{Key: "name", Value: request.Name},
		{Key: "email", Value: request.Email},
		{Key: "Password", Value: request.Password},
	}

	_, err := collection.InsertOne(*ctx, userData)
	if err != nil {
		return err
	}

	return nil
}
