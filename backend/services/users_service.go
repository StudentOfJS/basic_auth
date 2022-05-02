package services

import (
	"basic_auth/backend/domain/users"
	"basic_auth/backend/utils/errors"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	pwSlice, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, errors.NewBadRequestError("failed to encrypt the password")
	}
	user.Password = string(pwSlice[:])
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
