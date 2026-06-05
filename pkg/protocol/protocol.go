package protocol

type Cmd uint8

const (
	//public
	CmdLogin Cmd = iota + 1
	CmdRegister

	//for auth-ed
	CmdCheckPermission
	CmdGetMe
	CmdLogout

	//root
	CmdSetRole
	CmdGetAllUsers
	CmdDeleteUser
)
