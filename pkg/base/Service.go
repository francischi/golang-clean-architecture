package base

import (
	"golang/pkg/except"
)

type Service struct{	
}

func(s *Service) InvalidArgument(msg string)error{
	var InvalidArgument except.InvalidArgument
	InvalidArgument.Message = msg
	return &InvalidArgument
}

func(s *Service) SystemError(msg string)error{
	var SystemError except.SystemError
	SystemError.Message = msg
	return &SystemError
}