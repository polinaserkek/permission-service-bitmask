package tcp

import (
	"fmt"
	"net"
	"permission-service/internal/auth"
)

func CreateServer(service *auth.Service) {
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
		fmt.Println("tcp server...")

		//4
		go TcpHandler(conn, service)
	}
}
