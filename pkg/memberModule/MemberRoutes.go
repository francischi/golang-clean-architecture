package memberModule

import (
	"github.com/gin-gonic/gin"
	mw "golang/pkg/middleWare"
)

func SetRoute(g *gin.Engine ,baseGroup string){

	memberGroup :=g.Group(baseGroup+"/member")
	memberGroup.POST("" , InitMemberController().Create)
	memberGroup.POST("/login" , InitMemberController().LogIn)
	memberGroup.PATCH("/password" , mw.InitJwtMiddleWare().ConfirmToken , InitMemberController().ChangePwd)
}