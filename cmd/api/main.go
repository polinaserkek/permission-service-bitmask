package main

import (
	"fmt"
	"permission-service/db"
	"permission-service/internal/auth"
	"permission-service/internal/repository"
	"permission-service/internal/transport/tcp"
)

func main() {
	db, err := db.DatabaseConnect()

	if err != nil {
		fmt.Println("main.go says that there`s some err with db conn: ", err)
		return
	}

	userData := repository.NewUserRepository(db)

	service := auth.NewService(userData)

	tcp.CreateServer(service)

	// new login
	// loginedUser, err := service.Login("1bbc7dd3-3c89-47c3-ae0c-e4ea84d1d79c", )

	// new get
	// user, err := service.GetUser("1bbc7dd3-3c89-47c3-ae0c-e4ea84d1d79c")

	// if err != nil {
	// 	fmt.Println("user not found")

	// } else {
	// 	fmt.Println(user)
	// }

	//get all users

	// var users []*model.User
	// users, _ = service.GetAllUsers()
	// // users = append(users, users...)
	// for _, v := range users {
	// 	fmt.Println(v)
	// }

	// id := uuid.MustParse("f7155a86-1271-4124-ad1a-22aa852f8d6b")

	// err = service.SetRole(8, id, 3)
	// if err != nil {
	// 	fmt.Println("u r not root", err)
	// }

}
