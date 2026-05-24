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
	fmt.Println(user)
	fmt.Printf("%04b\n", user)
	user |= permission
	fmt.Println(user)
	fmt.Printf("%04b\n", user)
	return user

}

func HasPermission(user int, permission int) bool {
	if user&permission != 0 {
		fmt.Println("y")
		return true
	}
	fmt.Println("n")
	return false
}

func RemovePermission(user int, permission int) int {
	fmt.Println(user)
	user &^= permission
	// user = user &^ permission
	fmt.Println(user)

	return user
}
