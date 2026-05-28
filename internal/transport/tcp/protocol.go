package tcp

type Cmd uint8

const (
	CmdLogin Cmd = iota + 1
	CmdCheckPermission
	CmdGetRoles
)
