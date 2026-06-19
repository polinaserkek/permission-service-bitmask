package tcp

import (
	"net"
	"permission-service/internal/auth"
	"permission-service/internal/hash"
	"permission-service/internal/permissions"
	"permission-service/pkg/protocol"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func TcpHandler(conn net.Conn, service *auth.Service) error {
	buffer := make([]byte, 2048)

	n, err := conn.Read(buffer)

	if err != nil {
		return err
	}

	stringedResponse := string(buffer[:n])

	array := strings.Split(stringedResponse, "|")
	command := array[0]
	cmdInt, err := strconv.Atoi(command)

	if err != nil {
		return err
	}

	conn.Write([]byte(stringedResponse))
	defer conn.Close()

	switch protocol.Cmd(cmdInt) {

	case protocol.CmdLogin:
		username := array[1]
		pass := array[2]
		hashedPass, err := hash.HashPassword(hash.Password(pass))
		if err != nil {
			return err
		}
		service.Login(username, hash.Password(hashedPass))

	case protocol.CmdRegister:
		username := array[1]
		pass := array[2]

		hashedPass, err := hash.HashPassword(hash.Password(pass))
		if err != nil {
			return err
		}

		service.Register(username, hash.Password(hashedPass), permissions.Permission(1))

	case protocol.CmdCheckPermission:
		username := array[1]
		service.CmdCheckPermission(permissions.Permission(8), username)

	case protocol.CmdGetMe:
		id := array[1]
		service.GetUser(id)

	case protocol.CmdLogout:

	case protocol.CmdSetRole:
		id := array[1]
		parsedID := uuid.MustParse(id)
		role := array[2]
		roleInt, err := strconv.Atoi(role)

		if err != nil {
			return err
		}
		service.SetRole(permissions.Permission(8), parsedID, permissions.Permission(roleInt))

	case protocol.CmdGetAllUsers:
		service.GetAllUsers(permissions.Permission(8))

	case protocol.CmdDeleteUser:
		parsedID := uuid.MustParse("1bbc7dd3-3c89-47c3-ae0c-e4ea84d1d79c")
		service.DeleteUser(permissions.Permission(8), parsedID)
	}
	return nil
}
