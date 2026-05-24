package models

import "permission-service/internal/repository"

var Users = map[string]repository.Permission{
	"keker666": 3, //read + write
	"vasya":    7, // read...
	"poleno":   7, // read + write + delete
}
