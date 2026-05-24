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

	res := repository.HasPermission(users["vasya"], repository.Read)
	perm := repository.DecodePermissions(3)

	if res {
		fmt.Printf("Право %s есть у пользователя %s\n", perm, "vasya")
	}

	// repository.HasPermission(users["keker666"], repository.Delete)

	// users["keker666"] = repository.RemovePermission(users["keker666"], repository.Root)

}
