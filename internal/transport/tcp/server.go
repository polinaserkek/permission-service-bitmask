package tcp

import (
	"fmt"
	"net"
)

func CreateServer() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println(err)
			return
		}

		go TcpHandler(conn)
	}
}
