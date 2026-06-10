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
}
