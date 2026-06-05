package repository

import (
	"database/sql"
	"fmt"
	"permission-service/internal/model"
	"permission-service/internal/permissions"

	"github.com/google/uuid"
	_ "github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Login(id uuid.UUID, username string, password string) {

}

func (r *UserRepository) SetRole(id uuid.UUID, role permissions.Permission) error {
	query := `UPDATE users SET role=$2
	WHERE id=$1`
	fmt.Println("repo...")

	_, err := r.db.Exec(query, id, role)
	return err
}

func (r *UserRepository) CreateUser(
	id uuid.UUID,
	username string,
	password string,
	role permissions.Permission) (sql.Result, error) {

	query := `
	INSERT INTO users (id, username, password, role)
	VALUES
	($1, $2, $3, $4)
	`

	result, err := r.db.Exec(query, id, username, password, role)

	if err != nil {
		fmt.Println("Error to create user: ", err)

	}
	fmt.Println("repo:CreateUser()...")

	return result, err
}

func (r *UserRepository) GetUser(
	id uuid.UUID) (*model.User, error) {

	query := `
	SELECT id, username, password, role
	FROM users
	WHERE id = $1
	`
	var user model.User

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Role)

	if err != nil {
		fmt.Println("Error to get one user: ", err)
		return nil, err
	}
	return &user, err

}

func (r *UserRepository) GetAllUsers() ([]*model.User, error) {
	query := `SELECT * FROM users`
	users := []*model.User{}

	rows, err := r.db.Query(query)

	if err != nil {
		fmt.Println("Problem to get all users: ", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		user := model.User{}

		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Role,
		)

		if err != nil {
			fmt.Println("Problem to get all users: ", err)
			return nil, err
		}

		users = append(users, &user)
	}
	fmt.Println("repo: GetAllUsers()...")
	return users, err
}

func (r *UserRepository) DeleteUser(id uuid.UUID) error {
	query := "DELETE from users WHERE id=$1"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *UserRepository) UpdateUser(
	id uuid.UUID,
	username string,
	password string,
	role permissions.Permission) error {

	query := `UPDATE users SET
		username = $2,
		password = $3,
		role = $4
		WHERE id=$1`

	_, err := r.db.Exec(query, id, username, password, role)
	return err
}
