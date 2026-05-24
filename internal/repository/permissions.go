package repository

import "fmt"

const (
	Read   = 1 << iota //0001 -> 1
	Write              //0010 -> 2
	Delete             //0100 -> 4
	Root               //1000 -> 8

	// read + write -> 0011
	// write + delete -> 0110
	// read + delete -> 0101
	// read + write + delete -> 0111
	//....
)

func AddPermission(user int, permission int) int {
	// fmt.Println(user)
	// fmt.Printf("%04b\n", user)
	user |= permission
	// fmt.Println(user)
	// fmt.Printf("%04b\n", user)
	return user

}

func HasPermission(user int, permission int) (bool, string, string) {
	if user&permission != 0 {
		// fmt.Println(user)
		// fmt.Printf("%04b\n", user)
		// fmt.Println(user & permission)
		c := DecodePermission(user & permission)
		return true, c, "есть"
	} else {
		fmt.Println("from else...")
		c := DecodePermission(permission)
		return false, c, "нет"
	}
	// c := DecodePermission(permission)
	// return false, c, "нет"
}

func DecodePermission(permission int) string {
	fmt.Println(permission)
	switch permission {
	case 1:
		return "Read"
	case 2:
		return "Write"
	case 4:
		return "Delete"
	case 8:
		return "Root"
	}
	// fmt.Println(permission)
	// fmt.Printf("%04b\n", permission)
	return "aaaaa"
}

func RemovePermission(user int, permission int) int {
	fmt.Println(user)
	user &^= permission
	// user = user &^ permission
	fmt.Println(user)

	return user
}
