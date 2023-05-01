package dtos

import (
	"errors"
	"golang/pkg/helpers"
)

type ChangePwdDto struct {
	MemberId string
	OldPassword string
	NewPassword string
	ConfirmPassword string
}

func (dto *ChangePwdDto) Check()(error){
	if len(dto.MemberId) == 0 {
		return errors.New("member_id_required")
	}
	if len(dto.OldPassword) == 0 {
		return errors.New("old_password_required")
	}
	if len(dto.NewPassword) == 0 {
		return errors.New("new_password_required")
	}
	if !helpers.IsValidPassword(dto.NewPassword){
		return errors.New("invalid_password")
	}
	if len(dto.ConfirmPassword) == 0 {
		return errors.New("email_required")
	}
	if !helpers.IsValidPassword(dto.ConfirmPassword){
		return errors.New("invalid_password")
	}
	return nil
}