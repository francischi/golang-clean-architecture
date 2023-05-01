package helpers

import (
	"github.com/google/uuid"
)

func CreateUuid() string {
	return uuid.New().String()
}