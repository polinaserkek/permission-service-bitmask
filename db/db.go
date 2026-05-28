package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func DatabaseConnect(hostname string, port string, username string, password string, db string) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", hostname, port, username, password, db)
	conn, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println("From db/db.go: Error with opening db connection: ", err)
		return nil, err
	}

	err = conn.Ping()

	if err != nil {
		fmt.Println("From db/db.go: Error with pinging db: ", err)
		return nil, err
	}
	fmt.Println("connected to db!!!!!!!")
	return conn, nil
}
