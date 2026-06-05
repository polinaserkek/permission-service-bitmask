package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"permission-service/db"
	"permission-service/internal/auth"
	"permission-service/internal/repository"
	"permission-service/internal/transport/tcp"
)

func frontendTest() {
	conn, err := net.Dial("tcp", "localhost:8080")
	fmt.Println("main.go: frontendTest()")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	conn.Write([]byte("7"))
	io.Copy(os.Stdout, conn)
}

func main() {
	db, err := db.DatabaseConnect()

	if err != nil {
		fmt.Println("main.go says that there`s some err with db conn: ", err)
		return
	}

	//1
	userData := repository.NewUserRepository(db)

	//2
	service := auth.NewService(userData)

	//3
	go tcp.CreateServer(service)

	frontendTest()

	//new create!!!

	// _, err = service.Register("keker777", "123456", 4)

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
