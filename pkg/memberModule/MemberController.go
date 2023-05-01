package memberModule

import (
	// "fmt"
   "golang/pkg/memberModule/dtos"    
   "github.com/gin-gonic/gin"
   "golang/pkg/base"
)

type MemberController struct{
	MemberService MemberService
	base.Controller
}

func NewMemberController(memberService *MemberService) *MemberController{
	return &MemberController{
		MemberService : *memberService,
	}
}

func (c *MemberController)Create(g *gin.Context){

	var member dtos.CreateMemberDto
	g.Bind(&member)

	err := c.MemberService.Create(&member)
	if err!=nil{
		c.HandleError(g , err)
	}else {
		c.SuccessRes(g,nil)
	}
}

func (c *MemberController) LogIn(g *gin.Context){

	var member dtos.LogInDto
	g.Bind(&member)
	
	token , err :=c.MemberService.LogIn(&member)
	if err!=nil{
		c.HandleError(g , err)
	}else {
		c.SuccessRes(g,token)
	}
}

func (c *MemberController ) ChangePwd(g *gin.Context){
	
	var member dtos.ChangePwdDto
	g.Bind(&member)

	err := c.MemberService.ChangePwd(&member)
	if err!=nil{
		c.HandleError(g , err)
	}else {
		c.SuccessRes(g,"")
	}
}