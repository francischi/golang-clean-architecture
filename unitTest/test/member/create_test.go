package member

import (
	// "fmt"
	"golang/pkg/memberModule"
	"golang/pkg/memberModule/dtos"
	"golang/pkg/repos/models"
	"golang/unitTest/mockRepos"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/golang/mock/gomock"
)

func TestCreate(t *testing.T){
	controller := gomock.NewController(t)
    defer controller.Finish()
	
	memberRepo := mockRepos.NewMockMemberRepo(controller)

	memberRepo.EXPECT().IsEmailExist(gomock.Any()).Return(false).AnyTimes()

	memberRepo.EXPECT().Create(gomock.Any()).DoAndReturn(
		func(input models.MemberModel) interface{}{
			checked , errMsg := input.Check()
			if checked==false{
				return errMsg
			}
			return nil
		},
	).AnyTimes()
	
	memberService := memberModule.MemberService{
		MemberRepo: memberRepo,
	}

	t.Run("success", func(t *testing.T) {
		createDto := dtos.CreateMemberDto{
			Name : "frank",
			Email : "test@gmail.com",
			Gender : "FEMALE",
			Password : "thisisfrank",
		}

		err := memberService.Create(&createDto)
		assert.NoError(t, err)
	})

	t.Run("invalid name", func(t *testing.T) {

		createDto := dtos.CreateMemberDto{
			Email : "test@gmail.com",
			Gender : "FEMALE",
			Password : "thisisfrank",
		}

		err := memberService.Create(&createDto)
		assert.Error(t, err)
	})

	t.Run("required email", func(t *testing.T) {
		createDto := dtos.CreateMemberDto{
			Name : "frank",
			Gender : "FEMALE",
			Password : "thisisfrank",
		}

		err := memberService.Create(&createDto)
		assert.Error(t, err)
	})

	t.Run("invalid email", func(t *testing.T) {
		createDto := dtos.CreateMemberDto{
			Name : "frank",
			Email : "test.com",
			Gender : "FEMALE",
			Password : "thisisfrank",
		}

		err := memberService.Create(&createDto)
		assert.Error(t, err)
	})

	t.Run("invalid password", func(t *testing.T) {
		createDto := dtos.CreateMemberDto{
			Name : "frank",
			Email : "test@gmail.com",
			Password : "thisisfrank",
		}

		err := memberService.Create(&createDto)
		assert.Error(t, err)
	})
}