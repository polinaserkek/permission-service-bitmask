package permissions

type Permission uint8

const (
	Read Permission = 1 << iota
	Write
	Delete
	Root
)

func (p Permission) Add(permission Permission) Permission {
	p |= permission
	return p
}
func (p Permission) Has(permission Permission) bool {
	if p&permission != 0 {
		return true
	} else {
		return false
	}
}
func (p Permission) Remove(permission Permission) Permission {
	p &^= permission
	return p
}
func DecodePermissions(user Permission) []string {
	res := []string{}
	if user&Read != 0 {
		res = append(res, "Read")
	}
	if user&Write != 0 {
		res = append(res, "Write")
	}
	if user&Delete != 0 {
		res = append(res, "Delete")
	}
	if user&Root != 0 {
		res = append(res, "Root")
	}
	return res
}
