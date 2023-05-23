package memberModule

import (
	// "fmt"
	"errors"
	"golang/pkg/base"
	"golang/pkg/helpers"
	"golang/pkg/memberModule/dtos"
	impl "golang/pkg/repos/implement"
	"golang/pkg/repos/interfaces"
	"golang/pkg/repos/models"

	"gorm.io/gorm"
)

type MemberService struct {
	base.Service
	MemberRepo interfaces.MemberRepo
}

func NewMemberService(memberRepo *impl.MemberRepo) *MemberService{
	var MemberService MemberService
	MemberService.MemberRepo = memberRepo

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
		MemberId : helpers.CreateUuid(),
		Name : dto.Name,
		Gender: dto.Gender,
		Password : helpers.GetSHA256HashCode(dto.Password),
		Email: dto.Email,
		IsOnline: models.MEMBER_OFFLINE,
		CreateTime : helpers.GetTimeStamp(),
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

	hashedPassword := helpers.GetSHA256HashCode(dto.Password)
	if hashedPassword != memberModel.Password{
		return "",s.InvalidArgument("password_not_match")
	}
	var jwtToken helpers.JwtToken
	token,err = jwtToken.Create(memberModel.MemberId , memberModel.Name)
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
	
	if err!=nil{
		return s.InvalidArgument("member_repo_error")
	}
	if memberModel.Id==0{
		return s.InvalidArgument("no_matched_member_id")
	}

	if !s.confirmOldPwd(&dto.OldPassword ,&memberModel.Password){
		return s.InvalidArgument("old_password_not_match")
	}

	hashNewPwd := helpers.GetSHA256HashCode(dto.NewPassword) 
	
	if err := s.MemberRepo.ChangePwd(memberId ,hashNewPwd);err!=nil{
		return s.InvalidArgument("old_password_not_match")
	}
	return nil
}

func (s *MemberService) confirmOldPwd(pwd *string , oldPwd *string)(bool){
	hashPassword := helpers.GetSHA256HashCode(*pwd)
	return hashPassword == *oldPwd
}