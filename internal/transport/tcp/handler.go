package tcp

import (
	"fmt"
	"net"
	"permission-service/internal/auth"
	"permission-service/internal/permissions"
	"permission-service/pkg/protocol"
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

	array := strings.Split(stringedResponse, "|")

	fmt.Println(array)

	command := array[0]
	username := array[1]
	pass := array[2]
	role := array[3]

	cmdInt, err := strconv.Atoi(command)
	roleInt, err := strconv.Atoi(role)

	if err != nil {
		fmt.Println(err)
	}

	//6
	conn.Write([]byte(stringedResponse))
	defer conn.Close()

	///////////

	switch protocol.Cmd(cmdInt) {

	case protocol.CmdLogin:

	case protocol.CmdRegister:
		fmt.Println("handler: CmdRegister..")
		service.Register(username, pass, permissions.Permission(roleInt))

	case protocol.CmdCheckPermission:

	case protocol.CmdGetMe:

	case protocol.CmdLogout:

	case protocol.CmdSetRole:
		// parsedID := uuid.MustParse("1bbc7dd3-3c89-47c3-ae0c-e4ea84d1d79c")
		// service.SetRole(permissions.Permission(8), parsedID, permissions.Permission(roleInt))

	case protocol.CmdGetAllUsers:
		fmt.Println("CASE from handler...")
		service.GetAllUsers(permissions.Permission(8))

	case protocol.CmdDeleteUser:
		parsedID := uuid.MustParse("1bbc7dd3-3c89-47c3-ae0c-e4ea84d1d79c")
		service.DeleteUser(parsedID)
	}

}
