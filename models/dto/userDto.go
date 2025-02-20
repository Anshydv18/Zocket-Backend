package dto

import "github.com/golang-jwt/jwt/v5"

type UserDto struct {
	Email        string
	Name         string
	TaskAssigned interface{}
}

type AuthClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
