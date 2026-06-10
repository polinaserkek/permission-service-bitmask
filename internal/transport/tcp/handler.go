package tcp

import (
	"fmt"
	"net"
	"permission-service/internal/auth"
	"permission-service/internal/permissions"
	"permission-service/pkg/protocol"
	"strconv"
	"strings"
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

	array := strings.Split(stringedResponse, "|")

	// --------------------------------------------------
	// CmdLogin
	// command := array[0]
	// username := array[1]
	// pass := array[2]
	// hashedPass, err := hash.HashPassword(hash.Password(pass))

	// --------------------------------------------------
	// CmdRegister

	// command := array[0]
	// username := array[1]
	// pass := array[2]

	// hashedPass, err := hash.HashPassword(hash.Password(pass))
	// role := array[3]

	// cmdInt, err := strconv.Atoi(command)
	// roleInt, err := strconv.Atoi(role)

	// --------------------------------------------------
	// CmdGetMe

	// --------------------------------------------------
	// CmdLogout

	// --------------------------------------------------
	// CmdCheckPermission

	command := array[0]
	username := array[1]
	cmdInt, err := strconv.Atoi(command)

	// --------------------------------------------------
	// CmdSetRole

	// --------------------------------------------------
	// CmdGetAllUsers

	// --------------------------------------------------
	// CmdDeleteUser

	if err != nil {
		fmt.Println(err)
	}

	conn.Write([]byte(stringedResponse))
	defer conn.Close()

	///////////

	switch protocol.Cmd(cmdInt) {
	case protocol.CmdLogin:
		// service.Login(username, hash.Password(hashedPass))

	case protocol.CmdRegister:
		// service.Register(username, hash.Password(hashedPass), permissions.Permission(1))

	case protocol.CmdCheckPermission:
		service.CmdCheckPermission(permissions.Permission(8), username)

	case protocol.CmdGetMe:

	case protocol.CmdLogout:

	case protocol.CmdSetRole:
		// parsedID := uuid.MustParse("1bbc7dd3-3c89-47c3-ae0c-e4ea84d1d79c")
		// service.SetRole(permissions.Permission(8), parsedID, permissions.Permission(roleInt))

	case protocol.CmdGetAllUsers:
		// service.GetAllUsers(permissions.Permission(8))

	case protocol.CmdDeleteUser:
		// parsedID := uuid.MustParse("1bbc7dd3-3c89-47c3-ae0c-e4ea84d1d79c")
		// service.DeleteUser(permissions.Permission(roleInt), parsedID)
	}

}
