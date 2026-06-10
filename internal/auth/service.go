package auth

import (
	"database/sql"
	"errors"
	"fmt"

	"permission-service/internal/hash"
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

func (s *Service) Login(username string, password hash.Password) (bool, error) {
	foundPassword, err := s.repo.Login(username)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	result := hash.CheckPassword(password, foundPassword.Password)

	fmt.Printf("welcome, %s", username)
	return result, nil
}

func (s *Service) GetUser(id string) (*model.User, error) {
	parsedId := uuid.MustParse(id)
	fmt.Println("service GetUser()")
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
	for _, v := range users {
		fmt.Println(v)
	}
	return users, err
}

func (s *Service) Register(
	username string,
	password hash.Password,
	role permissions.Permission) (*sql.Result, error) {

	newId := uuid.New()
	result, _ := s.repo.CreateUser(newId, username, hash.Password(password), permissions.Permission(1))
	return &result, nil
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

func (s *Service) DeleteUser(caller permissions.Permission, id uuid.UUID) error {
	if caller.Has(permissions.Root) == false {
		return errors.New("forbidden")
	}
	err := s.repo.DeleteUser(id)
	return err
}

func (s *Service) CmdCheckPermission(
	caller permissions.Permission,
	username string) (permissions.Permission, error) {
	if caller.Has(permissions.Root) == false {
		return 0, errors.New("forbidden")
	}
	user, err := s.repo.Login(username)
	if err != nil {
		return 0, err
	}
	fmt.Println(permissions.DecodePermissions(user.Role))
	return user.Role, nil
}
