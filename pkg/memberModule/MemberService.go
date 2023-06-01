package memberModule

import (
	// "fmt"
	"errors"
	"golang/pkg/base"
	h "golang/pkg/helpers"
	"golang/pkg/memberModule/dtos"
	"golang/pkg/repos/interfaces"
	"golang/pkg/repos/models"
	"golang/pkg/tokenModule"

	"gorm.io/gorm"
)

type MemberService struct {
	base.Service
	MemberRepo interfaces.MemberRepo
	TokenService *tokenModule.TokenService
}

func NewMemberService(memberRepo interfaces.MemberRepo ,tokenService *tokenModule.TokenService) *MemberService{
	var MemberService MemberService
	MemberService.MemberRepo = memberRepo
	MemberService.TokenService = tokenService
	return &MemberService
}

func (s *MemberService)Create(dto *dtos.CreateMemberDto)(err error){
	if err := dto.Check();err!=nil{
		return s.InvalidArgument(err.Error())
	}
	if s.MemberRepo.IsEmailExist(dto.Email) {
		return s.InvalidArgument("account_existed")
	}
	MemberModel := models.MemberModel{
		MemberId : h.CreateUuid(),
		Name : dto.Name,
		Gender: dto.Gender,
		Password : h.GetSHA256HashCode(dto.Password),
		Email: dto.Email,
		IsOnline: models.MEMBER_OFFLINE,
		CreateTime : h.GetTimeStamp(),
	}
	error := s.MemberRepo.Create(MemberModel)

	return  error
}

func (s *MemberService) LogIn(dto *dtos.LogInDto)(token string ,err error){
	if err := dto.Check();err!=nil{
		return "",s.InvalidArgument(err.Error())
	}
	memberModel ,err:= s.MemberRepo.GetMemberByEmail(dto.Account)

	if errors.Is(err, gorm.ErrRecordNotFound){
		return "",s.InvalidArgument("no_matched_account")
	}

	hashedPassword := h.GetSHA256HashCode(dto.Password)
	if hashedPassword != memberModel.Password{
		return "",s.InvalidArgument("password_not_match")
	}
	token,err = s.TokenService.Create(memberModel.MemberId , memberModel.Name)
	if err != nil{
		return "",err
	}
	return token , nil
}

func (s *MemberService) ChangePwd(dto *dtos.ChangePwdDto)(err error){
	if err := dto.Check();err!=nil{
		return s.InvalidArgument(err.Error())
	}
	if dto.ConfirmPassword != dto.NewPassword {
		return s.InvalidArgument("confirm_password_not_match")
	}

	memberId := dto.MemberId
	memberModel ,err:= s.MemberRepo.GetMember(memberId)

	if errors.Is(err, gorm.ErrRecordNotFound){
		return s.InvalidArgument("no_matched_account")
	}
	if err!=nil{
		return s.SystemError("member_repo_error")
	}
	if memberModel.Id==0{
		return s.InvalidArgument("no_matched_account")
	}

	if !s.confirmOldPwd(&dto.OldPassword ,&memberModel.Password){
		return s.InvalidArgument("old_password_not_match")
	}

	hashNewPwd := h.GetSHA256HashCode(dto.NewPassword) 
	
	if err := s.MemberRepo.ChangePwd(memberId ,hashNewPwd);err!=nil{
		return s.SystemError("member_repo_error")
	}
	return nil
}

func (s *MemberService) confirmOldPwd(pwd *string , oldPwd *string)(bool){
	hashPassword := h.GetSHA256HashCode(*pwd)
	return hashPassword == *oldPwd
}