package member

// go get github.com/golang/mock/gomock
// go install github.com/golang/mock/mockgen
// mockgen -source="./pkg/repos/interfaces/MemberInterface.go" -destination="./unitTest/mockRepos/MemberRepoMock.go" -package="mockRepos"

import (
	// "fmt"
    "testing"
	"golang/pkg/memberModule"
	"golang/pkg/memberModule/dtos"
	"golang/pkg/repos/models"
    "golang/unitTest/mockRepos"
	"github.com/golang/mock/gomock"
)

func getMemberByEmailReturn()(member models.MemberModel){
    var model models.MemberModel
    model.Id = 1
    model.Email = "test@gmail.com"
    model.MemberId = "d69b16e9-61ee-46d1-a99e-894b760fa5a3"
    model.Name = "frank"
    model.CreateTime = 100
    model.DeleteTime = 100
    model.IsOnline = 0
    model.Password = "4cb6390de5467430a51932c0b87e8f1ec7ec33fbf08486f77704a6cd6f89c5de"
    
    return model
}

func logInDto()(dtos.LogInDto){
    var dto dtos.LogInDto
    dto.Account = "test@gmail.com"
    dto.Password = "thisisfrank4"
    return dto
}

func prepareMockRepo(controller *gomock.Controller) *mockRepos.MockMemberRepo {
    memberRepo := mockRepos.NewMockMemberRepo(controller)
    email := "test@gmail.com"
    memberModel  := getMemberByEmailReturn()
    memberRepo.EXPECT().GetMemberByEmail(email).Return(memberModel , nil)
    return memberRepo
}

func TestLogIn(t *testing.T){
    controller := gomock.NewController(t)
    defer controller.Finish()

    memberRepo := prepareMockRepo(controller)

    memberService := memberModule.MemberService{
        MemberRepo: memberRepo,
    }

    logInDto := logInDto()
    token , err := memberService.LogIn(&logInDto)
    if token =="" || err!=nil{
        t.Errorf("errMessage:%s",err.Error())
    }
}