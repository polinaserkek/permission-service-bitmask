package repository

import "permission-service/internal/auth"

// var Users = map[string]auth.Permission{"keker666": 3, "vasya": 7, "poleno": 7}

type UserRepository struct {
	Users map[string]auth.Permission
}

func (u UserRepository) GetUser(name string, permission auth.Permission) {

}
func (u UserRepository) GetAllUsers() {

}
func (u UserRepository) UpdateUser() {

}

func (u UserRepository) DeleteUser() {

}
func (u UserRepository) CreateUser() {

}
