package auth

type Permission uint8

const (
	//type Perm is inherited for Write, Del..
	Read Permission = 1 << iota
	Write
	Delete
	Root
)
