package main

import (
	"fmt"
	"permission-service/db"
	"permission-service/internal/auth"
	"permission-service/internal/model"
	"permission-service/internal/repository"
)

func main() {
	db, err := db.DatabaseConnect("localhost", "5432", "postgres", "12345", "postgres")

	if err != nil {
		fmt.Println("main.go says that there`s some err with db conn: ", err)
	}

	userData := repository.NewUserRepository(db)
	service := auth.NewService(userData)

	//new create!!!

	// _, err = service.Register("keker777", "123456", 4)

	// new login
	// loginedUsed, err := service.Login("1bbc7dd3-3c89-47c3-ae0c-e4ea84d1d79c", )

	//new get
	// _, err = service.GetUser("1bbc7dd3-3c89-47c3-ae0c-e4ea84d1d79c")

	// if err != nil {
	// 	fmt.Println("user not found")
	// } else {
	// 	fmt.Println("we ve found the user but we re too lazy to print him")
	// }

	//get all users

	var users []*model.User
	users, _ = service.GetAllUsers()
	// users = append(users, users...)
	for _, v := range users {
		fmt.Println(v)
	}

	// id := uuid.MustParse("f7155a86-1271-4124-ad1a-22aa852f8d6b")

	// err = service.SetRole(8, id, 3)
	// if err != nil {
	// 	fmt.Println("u r not root", err)
	// }

	// //UDPATE
	// id := uuid.MustParse("f7155a86-1271-4124-ad1a-22aa852f8d6b")
	// err = userData.UpdateUser(db, id, "test222", "1111", permissions.Delete)
	// if err != nil {
	// 	fmt.Println("Problem to get all users: ", err)
	// }

	// //GET ALL  !!!!!!!!!!!!!!!!!!!!

	// users, err := userData.GetAllUsers()
	// if err != nil {
	// 	fmt.Println("Problem to get all users: ", err)
	// }
	// fmt.Println(users)

	// for _, v := range users {
	// 	fmt.Println(v)
	// }

	// DELETE
	// id := uuid.MustParse("2f948611-e89e-44a4-bf1c-b376c837f078")
	// err = userType.DeleteUser(id)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	////////

	/*roles*/
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
