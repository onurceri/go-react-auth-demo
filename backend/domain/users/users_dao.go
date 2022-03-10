package users

import (
	"github.com/onurceri/go-react-auth-demo/backend/datasource/mysql/users_db"
	"github.com/onurceri/go-react-auth-demo/backend/utils/errors"
)

var (
	queryInsertUser     = "INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?);"
	queryGetUserByEmail = "SELECT id, first_name, last_name, email, password FROM users WHERE email = ?"
	queryGetUserById    = "SELECT id, first_name, last_name, email FROM users WHERE id = ?"
)

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewBadRequestError("database error")
	}

	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password)
	if saveErr != nil {
		return errors.NewInternalServerError("database error")
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	user.ID = userID
	return nil
}

func (user *User) GetByEmail() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUserByEmail)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Email)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}

	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password); err != nil {
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) GetUserByID() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUserById)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}

	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); err != nil {
		return errors.NewInternalServerError("database error")
	}

	return nil
}
