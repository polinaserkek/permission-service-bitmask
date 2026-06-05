package main

import (
	"fmt"
	"net"
	"permission-service/pkg/protocol"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("conn err on the client", err)
	}

	defer conn.Close()

}

func SendCommand(cmd protocol.Cmd) {

}
