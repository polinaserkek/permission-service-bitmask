package auth

import (
	"database/sql"
	"errors"
	"fmt"

	"permission-service/internal/auth/errs"
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
		if errors.Is(err, sql.ErrNoRows) {
			return false, errs.ErrUserNotFound
		}
		return false, err
	}
	result := hash.CheckPassword(password, foundPassword.Password)
	if !result {
		return false, errs.ErrIncorrectPassword
	}

	fmt.Printf("welcome, %s", username)
	return result, nil
}

func (s *Service) GetUser(id string) (*model.User, error) {
	parsedId := uuid.MustParse(id)
	foundUser, err := s.repo.GetUser(parsedId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrUserNotFound
		}
		return nil, err
	}
	return foundUser, nil
}

func (s *Service) GetAllUsers(caller permissions.Permission) ([]*model.User, error) {
	if caller.Has(permissions.Root) == false {
		return nil, errs.ErrNotRoot
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
	result, err := s.repo.CreateUser(newId, username, hash.Password(password), permissions.Permission(1))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *Service) SetRole(
	caller permissions.Permission,
	targetID uuid.UUID,
	role permissions.Permission) error {

	if caller.Has(permissions.Root) == false {
		return errs.ErrNotRoot
	}
	return s.repo.SetRole(targetID, role)
}

func (s *Service) DeleteUser(caller permissions.Permission, id uuid.UUID) error {
	if caller.Has(permissions.Root) == false {
		return errs.ErrNotRoot
	}
	err := s.repo.DeleteUser(id)
	return err
}

func (s *Service) CmdCheckPermission(
	caller permissions.Permission,
	username string) (permissions.Permission, error) {
	if caller.Has(permissions.Root) == false {
		return 0, errs.ErrNotRoot
	}
	user, err := s.repo.Login(username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errs.ErrUserNotFound
		}
		return 0, err
	}
	fmt.Println(permissions.DecodePermissions(user.Role))
	return user.Role, nil
}
