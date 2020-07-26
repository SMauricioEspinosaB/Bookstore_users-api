package users

import (
	"fmt"

	"github.com/SMauricioEspinosaB/Bookstore_users-api/datasources/mysql/users_db"
	"github.com/SMauricioEspinosaB/Bookstore_users-api/utils/errors"
	"github.com/SMauricioEspinosaB/Bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name,last_name,email,date_created,status, password) VALUES (?,?,?,?,?,?)"
	queryGetUser          = "SELECT id,first_name,last_name,email,date_created, status FROM users WHERE id=?;"
	queryUpdateUser       = "UPDATE users SET first_name = ?,last_name = ?,email = ? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id =?"
	queryFindUserByStatus = "SELECT id,first_name,last_name,email,date_created,status FROM users WHERE status =?"
)

/*
var (
	usersDB = make(map[int64]*User)
)
*/

func something() {
	user := User{}
	if err := user.Get(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user.FirstName)
}

func (user *User) Get() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	// single row
	result := stmt.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}

	return nil
	//multiples user
	/*
		results, _ := stmt.Query(user.Id)
		defer results.Close()
	*/

	/*
		result := usersDB[user.Id]
		if result == nil {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
		}

		user.Id = result.Id
		user.FirstName = result.FirstName
		user.LastName = result.LastName
		user.Email = result.Email
		user.DateCreated = result.DateCreated

		return nil
	*/
}

func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {

		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	inserResult, SaveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if SaveErr != nil {
		return mysql_utils.ParseError(SaveErr)
	}

	userID, err := inserResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(SaveErr)
	}
	user.Id = userID
	return nil

	//	result, err := users_db.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)
	/*
		current := usersDB[user.Id]
		if current != nil {
			if current.Email == user.Email {
				return errors.NewBadRequestError(fmt.Sprintf("user %s already registered", user.Email))
			}
			return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
		}

		user.DateCreated = date_utils.GetNowString()
		usersDB[user.Id] = user
	*/
}
func (user *User) Update() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil

}
