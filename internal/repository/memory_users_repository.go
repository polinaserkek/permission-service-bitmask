package repository

import (
	"permission-service/internal/auth"
)

// отдельный тип    не объект!!
type UserRepository struct {
	Users map[string]auth.Permission
}

func (u *UserRepository) GetUser(name string) auth.Permission {
	return u.Users[name]
}

func (u *UserRepository) DeleteUser(name string) {
	delete(u.Users, name)
}

func (u *UserRepository) UpdateUser(name string, permission auth.Permission) {
	u.Users[name] = permission
}

func (u *UserRepository) CreateUser(name string, permission auth.Permission) {
	u.Users[name] = permission
}
