package protocol

type Cmd uint8

const (
	//public
	CmdLogin Cmd = iota + 1
	CmdRegister

	//for auth-ed
	CmdGetMe
	CmdLogout

	//root
	CmdCheckPermission
	CmdSetRole
	CmdGetAllUsers
	CmdDeleteUser
)

func CodeProtocol(role Cmd) {

}
