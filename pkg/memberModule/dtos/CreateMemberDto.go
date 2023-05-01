package dtos

import (
	"errors"
	"golang/pkg/helpers"
	"golang/pkg/repos/models"
)

type CreateMemberDto struct {
	Name     string
	Gender   string  
	Password string  
	Email    string 
}

func (member *CreateMemberDto) Check() (error){
	if len(member.Name) == 0 {
		return errors.New("name_required")
	}
	if len(member.Email) == 0 {
		return errors.New("email_required")
	}
	if !helpers.IsValidEmail(member.Email){
		return errors.New("invalid_email")
	}
	if len(member.Gender) == 0 {
		return errors.New("gender_required")
	}
	if member.Gender != models.MEMBER_GENDER_FEMALE && member.Gender !=  models.MEMBER_GENDER_MALE{
		return errors.New("invalid_gender")
	}
	if len(member.Password) == 0 {
		return errors.New("password_required")
	}
	if !helpers.IsValidPassword(member.Password){
		return errors.New("invalid_password")
	}
	return nil
}