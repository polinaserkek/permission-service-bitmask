package auth

import (
	"database/sql"
	"errors"
	"fmt"

	// "permission-service/internal/auth"

	"permission-service/internal/model"
	"permission-service/internal/permissions"
	"permission-service/internal/repository"

	"github.com/google/uuid"
)

type Service struct {
	repo *repository.UserRepository
}

func NewService(repo *repository.UserRepository) *Service {
	return &Service{repo: repo}
}

// must be corrected!!!
// func (s *Service) Login(id string, username string, password string) (*model.User, error) {
// 	parsedId := uuid.MustParse(id)
// 	foundUser, err := s.repo.Login(parsedId)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return foundUser, nil
// }

func (s *Service) GetUser(id string) (*model.User, error) {
	parsedId := uuid.MustParse(id)
	foundUser, err := s.repo.GetUser(parsedId)
	if err != nil {
		return nil, err
	}
	return foundUser, nil

}

func (s *Service) GetAllUsers(caller permissions.Permission) ([]*model.User, error) {
	if caller.Has(permissions.Root) == false {
		return nil, errors.New("forbidden")
	}
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	fmt.Println("service.go: GetAllUsers()...")
	for _, v := range users {
		fmt.Println(v)
	}
	return users, err
}

func (s *Service) Register(

	username string,
	password string,
	role permissions.Permission) (*sql.Result, error) {

	newId := uuid.New()
	result, _ := s.repo.CreateUser(newId, username, password, role)
	return &result, nil
	// userType.CreateUser(db, newId, "test", "12345", 6)
}

func (s *Service) SetRole(
	caller permissions.Permission,
	targetID uuid.UUID,
	role permissions.Permission) error {

	if caller.Has(permissions.Root) == false {
		return errors.New("forbidden")
	}
	return s.repo.SetRole(targetID, role)
}
