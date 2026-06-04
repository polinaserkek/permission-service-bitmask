package tcp

import (
	"fmt"
	"net"
	"permission-service/internal/auth"
	"permission-service/internal/permissions"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func TcpHandler(conn net.Conn, service *auth.Service) {
	buffer := make([]byte, 2048)
	n, err := conn.Read(buffer)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("tcp handler.....")

	stringedResponse := string(buffer[:n])

	res2 := strings.Split(stringedResponse, "|")
	elem := res2[0]
	cmdInt, err := strconv.Atoi(elem)

	conn.Write([]byte(stringedResponse))
	conn.Close()

	switch Cmd(cmdInt) {

	case CmdLogin:

	case CmdRegister:
		service.Register("lol", "11111", permissions.Permission(3))

	case CmdCheckPermission:

	case CmdGetMe:

	case CmdLogout:

	case CmdSetRole:
		parsedID := uuid.MustParse("1bbc7dd3-3c89-47c3-ae0c-e4ea84d1d79c")
		service.SetRole(permissions.Permission(8), parsedID, permissions.Permission(8))

	case CmdGetAllUsers:

	case CmdDeleteUser:

	}
}
