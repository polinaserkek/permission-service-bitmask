package main

import (
	"fmt"
	"permission-service/internal/models"
	"permission-service/internal/repository"
)

func main() {
	users := models.Users

	// users["keker666"] = repository.AddPermission(users["keker666"], repository.Delete)
	// users["keker666"] = repository.AddPermission(users["keker666"], repository.Root)

	_, s, m := repository.HasPermission(users["vasya"],
		repository.Root)

	if s != "" {
		fmt.Printf("Право %s %s у пользователя %s\n", s, m, "vasya")
	}

	// repository.HasPermission(users["keker666"], repository.Delete)

	// users["keker666"] = repository.RemovePermission(users["keker666"], repository.Root)

}
