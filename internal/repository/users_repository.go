package repository

import (
	"database/sql"
	"permission-service/internal/hash"
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

func (r *UserRepository) Login(username string) (*model.User, error) {
	query := "SELECT password, role FROM users WHERE username=$1"
	var user model.User
	row := r.db.QueryRow(query, username)
	err := row.Scan(
		&user.Password,
		&user.Role,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) SetRole(id uuid.UUID, role permissions.Permission) error {
	query := `UPDATE users SET role=$2
	WHERE id=$1`

	_, err := r.db.Exec(query, id, role)
	return err
}

func (r *UserRepository) CreateUser(
	id uuid.UUID,
	username string,
	password hash.Password,
	role permissions.Permission) (sql.Result, error) {

	query := `
	INSERT INTO users (id, username, password, role)
	VALUES
	($1, $2, $3, $4)
	`

	result, err := r.db.Exec(query, id, username, password, role)

	if err != nil {
		return nil, err
	}

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

	row := r.db.QueryRow(query, id)
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Role)

	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (r *UserRepository) GetAllUsers() ([]*model.User, error) {
	query := `SELECT * FROM users`
	users := []*model.User{}

	rows, err := r.db.Query(query)

	if err != nil {
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
			return nil, err
		}

		users = append(users, &user)
	}
	return users, nil
}

func (r *UserRepository) DeleteUser(id uuid.UUID) error {
	query := "DELETE from users WHERE id=$1"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *UserRepository) UpdateUser(
	id uuid.UUID,
	username string,
	password hash.Password,
	role permissions.Permission) error {

	query := `UPDATE users SET
		username = $2,
		password = $3,
		role = $4
		WHERE id=$1`

	_, err := r.db.Exec(query, id, username, password, role)
	return err
}
