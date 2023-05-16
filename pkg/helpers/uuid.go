package helpers

import (
	"github.com/google/uuid"
)

func CreateUuid() string {
	return uuid.New().String()
}

func CheckUuid(id string) bool {
	_,err := uuid.Parse(id)
	if err == nil{
		return true
	}
	return false
}