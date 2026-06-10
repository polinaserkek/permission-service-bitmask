package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"permission-service/pkg/protocol"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("conn err on the client", err)
	}

	defer conn.Close()

	cmd := "1|test222|1111"

	conn.Write([]byte(cmd))
	io.Copy(os.Stdout, conn)

}

func SendCommand(cmd protocol.Cmd) {

}
