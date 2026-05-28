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

	/////////

	/*roles*/
	kekerPerms := repo.GetUser("keker666")
	fmt.Println(kekerPerms)
	//добавили рута
	kekerPerms = kekerPerms.Add(auth.Root)
	fmt.Println(kekerPerms)

	//
	repo.UpdateUser("keker666", kekerPerms)

	//убрали рута
	kekerPerms = kekerPerms.Remove(auth.Root)

	//////

	vasyaPerms := repo.GetUser("vasya")

	res := vasyaPerms.Has(auth.Read)
	if res {
		fmt.Printf("Право %s есть у пользователя %s\n", "Read", "vasya")
	} else {
		fmt.Println("Нет права")
	}

	permList := auth.DecodePermissions(vasyaPerms)
	fmt.Printf("Все права пользователя %s: %v", "vasya", permList)

}
