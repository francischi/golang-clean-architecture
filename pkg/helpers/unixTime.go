package helpers

import (
	"time"
)

func GetTimeStamp() int {
	return int(time.Now().Unix())
}