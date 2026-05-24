package repository

type Permission int

const (
	//type Perm is inherited for Write, Del..
	Read   Permission = 1 << iota //0001 -> 1
	Write                         //0010 -> 2
	Delete                        //0100 -> 4
	Root                          //1000 -> 8

	// read + write -> 0011
	// write + delete -> 0110
	// read + delete -> 0101
	// read + write + delete -> 0111
	//....
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
	/* 0111  0001   */
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
