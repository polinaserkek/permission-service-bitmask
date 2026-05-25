package main

import (
	"fmt"
	"permission-service/internal/auth"
	"permission-service/internal/repository"
)

func main() {
	users := repository.Users

	users["keker666"] = users["keker666"].Add(auth.Root)

	res := users["vasya"].Has(auth.Read)
	if res {
		fmt.Printf("Право %s есть у пользователя %s\n", "Read", "vasya")
	} else {
		fmt.Println("Нет права")
	}

	perm := auth.DecodePermissions(users["vasya"])
	fmt.Printf("Все права пользователя %s: %v", "vasya", perm)

	users["keker666"] = users["keker666"].Remove(auth.Root)

}
