package db

import (
	"database/sql"
	"fmt"
	"os"
	"permission-service/internal/config"

	_ "github.com/lib/pq"
)

func DatabaseConnect() (*sql.DB, error) {
	config.LoadEnv()
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	// connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", hostname, port, username, password, db)
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
