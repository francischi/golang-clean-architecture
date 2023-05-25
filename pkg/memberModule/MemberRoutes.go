package memberModule

import (
	"github.com/gin-gonic/gin"
	mw "golang/pkg/middleWare"
)

func SetRoute(g *gin.Engine ,baseGroup string){

	memberGroup :=g.Group(baseGroup+"/member")
	memberGroup.POST("" , 
		func(ctx *gin.Context){InitMemberController().Create(ctx)},
	)
	memberGroup.POST("/login" ,
		func(ctx *gin.Context){InitMemberController().LogIn(ctx)},
	)
	memberGroup.PATCH("/password" , 
		func(ctx *gin.Context){mw.InitJwtMiddleWare().ConfirmToken(ctx)} , 
		func(ctx *gin.Context){InitMemberController().ChangePwd(ctx)},
	)
}