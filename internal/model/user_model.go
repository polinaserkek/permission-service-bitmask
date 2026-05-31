package model

import (
	"permission-service/internal/permissions"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Username string
	Password string
	Role     permissions.Permission
}
