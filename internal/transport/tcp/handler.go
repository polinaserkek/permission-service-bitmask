package tcp

import (
	"fmt"
	"net"
)

func TcpHandler(conn net.Conn) {
	buffer := make([]byte, 2048)
	n, err := conn.Read(buffer)

	if err != nil {
		fmt.Println(err)
		return
	}

	stringedResponse := string(buffer[:n])
	fmt.Println(stringedResponse)

	// switch stringedResponse {
	// 	case CmdSetRole: auth.Service.SetRole()
	// }
}
