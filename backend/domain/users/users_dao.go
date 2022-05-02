package users

import (
	"basic_auth/backend/datasource/mysql/users_db"
	"basic_auth/backend/utils/errors"
)

var (
	queryInsertUser = "INSERT INTO users (first__name, last_name, email, password) VALUES (?, ?, ?, ?);"
)

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError("database error")
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