package main

import (
	"permission-service/internal/models"
	"permission-service/internal/repository"
)

func main() {
	users := models.Users

	add := repository.AddPermission(users["keker666"], repository.Delete)
	add()
	add = repository.AddPermission(users["keker666"], repository.Root)
	add()

	// repository.HasPermission(users["keker666"], repository.Delete)
}
