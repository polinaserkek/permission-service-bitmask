package model

import (
	"permission-service/internal/hash"
	"permission-service/internal/permissions"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Username string
	Password hash.Password
	Role     permissions.Permission
}
