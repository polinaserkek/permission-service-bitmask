package model

import (
	"permission-service/internal/auth"

	"github.com/google/uuid"
)

type User struct {
	id       uuid.UUID
	username string
	password string
	role     auth.Permission
}
