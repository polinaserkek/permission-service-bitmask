package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("conn err on the client", err)
	}

	defer conn.Close()
	SendCommand()

	cmd := "3|8665e896-26d8-4513-9053-4c06f992c264|7"

	conn.Write([]byte(cmd))
	io.Copy(os.Stdout, conn)

}

func SendCommand(sentData string) {

}
