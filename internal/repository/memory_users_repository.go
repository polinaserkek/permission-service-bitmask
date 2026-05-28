package repository

import (
	"database/sql"
	"fmt"
	"permission-service/internal/auth"

	"github.com/google/uuid"
	_ "github.com/google/uuid"
)

// // отдельный тип    не объект!!
// type UserRepository struct {
// 	Users map[string]auth.Permission
// 	// id uuid
// }

type User struct {
	id       uuid.UUID
	username string
	password string
	role     auth.Permission
}

func (u *User) CreateUser(
	dbConn *sql.DB,
	id uuid.UUID,
	username string,
	password string,
	role auth.Permission) (sql.Result, error) {

	query := `
	INSERT INTO users (id, username, password, role)
	VALUES
	($1, $2, $3, $4)
	`

	result, err := dbConn.Exec(query, id, username, password, role)

	if err != nil {
		fmt.Println("Error to create user: ", err)

	}

	return result, err
}

func (u *User) GetUser(
	dbConn *sql.DB,
	id uuid.UUID) (*User, error) {

	query := `
	SELECT id,
	username,
	password,
	role
	FROM users
	WHERE id = $1
	`
	var user User

	err := dbConn.QueryRow(query, id).Scan(
		&user.id,
		&user.username,
		&user.password,
		&user.role)

	if err != nil {
		fmt.Println("Error to get one user: ", err)
		return nil, err
	}
	return &user, err

}

func (u *User) GetAllUsers(dbConn *sql.DB) ([]*User, error) {
	query := `SELECT * FROM users`
	users := []*User{}

	rows, err := dbConn.Query(query)

	if err != nil {
		fmt.Println("Problem to get all users: ", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		user := User{}

		err := rows.Scan(
			&user.id,
			&user.username,
			&user.password,
			&user.role,
		)

		if err != nil {
			fmt.Println("Problem to get all users: ", err)
			return nil, err
		}

		users = append(users, &user)
	}
	return users, err
}

func (u *User) DeleteUser(name string) {
}

func (u *User) UpdateUser(name string, permission auth.Permission) {

}
