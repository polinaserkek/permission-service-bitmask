package main

import (
	"fmt"
	"permission-service/db"
	"permission-service/internal/repository"

	"github.com/google/uuid"
)

func main() {
	db, err := db.DatabaseConnect("localhost", "5432", "postgres", "12345", "postgres")

	if err != nil {
		fmt.Println("main.go says that there`s some err with db conn: ", err)
	}

	userType := repository.User{}

	// //CREATE  !!!!!!!!!!!!!!!!!!!!
	// newId := uuid.New()
	// userType.CreateUser(db, newId, "test", "12345", 6)

	//GET ONE  !!!!!!!!!!!!!!!!!!!!
	// id := uuid.MustParse("2f948611-e89e-44a4-bf1c-b376c837f078")
	// poleno, err := userType.GetUser(db, id)

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(poleno)

	//UDPATE
	id := uuid.MustParse("f7155a86-1271-4124-ad1a-22aa852f8d6b")
	err = userType.UpdateUser(db, id, "test222", "1111", 3)
	if err != nil {
		fmt.Println("Problem to get all users: ", err)
	}
	//GET ALL  !!!!!!!!!!!!!!!!!!!!

	users, err := userType.GetAllUsers(db)
	if err != nil {
		fmt.Println("Problem to get all users: ", err)
	}
	fmt.Println(users)

	for _, v := range users {
		fmt.Println(v)
	}

	// DELETE
	// id := uuid.MustParse("2f948611-e89e-44a4-bf1c-b376c837f078")
	// err = userType.DeleteUser(db, id)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	////////

	// repo := repository.UserRepository{
	// 	Users: map[string]auth.Permission{
	// 		"keker666": 3,
	// 		"vasya":    7,
	// 		"poleno":   7,
	// 	},
	// }

	// /*roles*/
	// kekerPerms := repo.GetUser("keker666")
	// fmt.Println(kekerPerms)
	// //добавили рута
	// kekerPerms = kekerPerms.Add(auth.Root)
	// fmt.Println(kekerPerms)

	// //
	// repo.UpdateUser("keker666", kekerPerms)

	// //убрали рута
	// kekerPerms = kekerPerms.Remove(auth.Root)

	// //////

	// vasyaPerms := repo.GetUser("vasya")

	// res := vasyaPerms.Has(auth.Read)
	// if res {
	// 	fmt.Printf("Право %s есть у пользователя %s\n", "Read", "vasya")
	// } else {
	// 	fmt.Println("Нет права")
	// }

	// permList := auth.DecodePermissions(vasyaPerms)
	// fmt.Printf("Все права пользователя %s: %v", "vasya", permList)

}
