package member

import (
	// "fmt"
	"golang/pkg/memberModule"
	"golang/pkg/memberModule/dtos"
	"golang/pkg/repos/models"
	"golang/unitTest/mockRepos"
	"testing"

	"github.com/golang/mock/gomock"
)

func createDto()(dtos.CreateMemberDto){
	var dto dtos.CreateMemberDto
	dto.Name = "frank"
	dto.Email = "test@gmail.com"
	dto.Gender = "FEMALE"
	dto.Password = "thisisfrank"
	return dto
}

func createInput()(member models.MemberModel){
	var model models.MemberModel
	return model
}

func TestCreate(t *testing.T){
	controller := gomock.NewController(t)
    defer controller.Finish()
	
	memberRepo := mockRepos.NewMockMemberRepo(controller)
	
	memberRepo.EXPECT().IsEmailExist(gomock.Any()).Return(false)

	memberRepo.EXPECT().Create(gomock.Any()).DoAndReturn(
		func(input models.MemberModel) interface{}{
			checked , errMsg := input.Check()
			if checked==false{
				return errMsg
			}
			return nil
		},
	)

	memberService := memberModule.MemberService{
        MemberRepo: memberRepo,
    }

	createDto := createDto()

	err := memberService.Create(&createDto)
	if err!= nil{
		t.Errorf("errMessage:%s",err.Error())
	}
}