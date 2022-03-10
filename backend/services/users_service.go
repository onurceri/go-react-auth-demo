package services

import (
	"github.com/onurceri/go-react-auth-demo/backend/domain/users"
	"github.com/onurceri/go-react-auth-demo/backend/utils/errors"
	"golang.org/x/crypto/bcrypt"
	"github.com/onurceri/go-react-auth-demo/backend/datasource/mysql/users_db"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	pwSlice, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, errors.NewBadRequestError("failed to encrypt password")
	}
	user.Password = string(pwSlice[:])

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(user users.User) (*users.User, *errors.RestErr) {
	result := &users.User{Email: user.Email}

	if err := result.GetByEmail(); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
		return nil, errors.NewBadRequestError("invalid password")
	}

	resultWp := &users.User{ID: result.ID, FirstName: result.FirstName, LastName: result.LastName, Email: result.Email}

	return resultWp, nil
}

func GetUserByID(userId int64) (*users.User, *errors.RestErr) {
	err := users_db.Client.Ping()
	if err != nil {
		println(err)
	}
	if users_db.Client == nil {
		print("alksdjhasd")
	}
	result := &users.User{ID: userId}

	if err := result.GetUserByID(); err != nil {
		return nil, err
	}

	return result, nil
}
