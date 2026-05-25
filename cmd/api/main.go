package main

import (
	"fmt"
	"permission-service/internal/auth"
	"permission-service/internal/repository"
)

func main() {

	repo := repository.UserRepository{
		Users: map[string]auth.Permission{
			"keker666": 3,
			"vasya":    7,
			"poleno":   7,
		},
	}

	repo.Users["keker666"] = repo.Users["keker666"].Add(auth.Root)

	res := repo.Users["vasya"].Has(auth.Read)
	if res {
		fmt.Printf("Право %s есть у пользователя %s\n", "Read", "vasya")
	} else {
		fmt.Println("Нет права")
	}

	perm := auth.DecodePermissions(repo.Users["vasya"])
	fmt.Printf("Все права пользователя %s: %v", "vasya", perm)

	repo.Users["keker666"] = repo.Users["keker666"].Remove(auth.Root)

}
