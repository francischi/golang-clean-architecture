package dtos

import (
	// "fmt"
	"errors"
	"golang/pkg/helpers"
)

type LogInDto struct{
	Account string
	Password string
}

func(dto *LogInDto) Check()(error){
	if len(dto.Account) == 0 {
		return errors.New("account_required")
	}
	if !helpers.IsValidEmail(dto.Account){
		return errors.New("invalid_account")
	}
	if len(dto.Password) == 0 {
		return errors.New("password_required")
	}
	if !helpers.IsValidPassword(dto.Password){
		return errors.New("invalid_password")
	}
	return nil
}