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
	if res {
		fmt.Printf("Право %s есть у пользователя %s\n", "Read", "vasya")
	}

	perm := repository.DecodePermissions(users["vasya"])
	fmt.Printf("Все права пользователя %s: %v", "vasya", perm)

	// repository.HasPermission(users["keker666"], repository.Delete)

	// users["keker666"] = repository.RemovePermission(users["keker666"], repository.Root)

}
