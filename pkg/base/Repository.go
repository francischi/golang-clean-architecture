package base

import (
	// "fmt"
	// "errors"
	"golang/pkg/except"
)

type Repository struct{
}

func(r *Repository) InvalidArgument(msg string)error{
	var InvalidArgument except.InvalidArgument
	InvalidArgument.Message = msg
	return &InvalidArgument
}

func(r *Repository) SystemError(msg string)error{
	var SystemError except.SystemError
	SystemError.Message = msg
	return &SystemError
}