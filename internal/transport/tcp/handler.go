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

	//5
	n, err := conn.Read(buffer)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("tcp handler.....")

	stringedResponse := string(buffer[:n])

	res2 := strings.Split(stringedResponse, "|")
	fmt.Println(res2)
	command := res2[0]
	// username := res2[1]
	// pass := res2[2]
	// role := res2[2]
	cmdInt, err := strconv.Atoi(command)
	fmt.Println(cmdInt)
	// roleInt, err := strconv.Atoi(role)

	if err != nil {
		fmt.Println(err)
	}

	//6
	conn.Write([]byte(stringedResponse))
	conn.Close()

	///////////

	switch Cmd(cmdInt) {

	case CmdLogin:

	case CmdRegister:
		// service.Register(username, pass, permissions.Permission(roleInt))

	case CmdCheckPermission:

	case CmdGetMe:

	case CmdLogout:

	case CmdSetRole:
		// parsedID := uuid.MustParse("1bbc7dd3-3c89-47c3-ae0c-e4ea84d1d79c")
		// service.SetRole(permissions.Permission(8), parsedID, permissions.Permission(roleInt))

	case CmdGetAllUsers:
		fmt.Println("CASE from handler...")
		service.GetAllUsers(permissions.Permission(8))

	case CmdDeleteUser:
		parsedID := uuid.MustParse("1bbc7dd3-3c89-47c3-ae0c-e4ea84d1d79c")
		service.DeleteUser(parsedID)
	}
}
