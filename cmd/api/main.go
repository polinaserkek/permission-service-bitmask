package main

import (
	"permission-service/internal/models"
	"permission-service/internal/repository"
)

func main() {
	users := models.Users

	users["keker666"] = repository.AddPermission(users["keker666"], repository.Delete)
	users["keker666"] = repository.AddPermission(users["keker666"], repository.Root)

	// repository.HasPermission(users["keker666"], repository.Delete)

	users["keker666"] = repository.RemovePermission(users["keker666"], repository.Root)
}
