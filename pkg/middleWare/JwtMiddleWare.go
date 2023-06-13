package middleWare

import (
	// "fmt"
	"golang/pkg/base"
	"github.com/gin-gonic/gin"
	"golang/pkg/tokenModule"
)

func InitJwtMiddleWare()(middleWare *JwtMiddleWare){

	var JwtMiddleWare JwtMiddleWare
 	return &JwtMiddleWare
}

type JwtMiddleWare struct {
	base.MiddleWare
	tokenService tokenModule.TokenService
}

func (m *JwtMiddleWare) ConfirmToken(g *gin.Context){
	header := g.Request.Header["Bearer-Token"]
	if len(header)==0{
		m.InvaliAugument(g,"token_required")
		return
	}
	jwtToken := g.Request.Header["Bearer-Token"][0]

	val,err  := m.tokenService.IsValidJwt(jwtToken)
	if err!=nil{
		m.SystemError(g,err.Error())
		return
	}
	if !val {
		m.InvaliAugument(g,"invalid_token")
		return
	}

	val,err  = m.tokenService.IsJwtInTime(jwtToken)
	if err!=nil{
		m.SystemError(g,err.Error())
		return
	}
	if !val {
		m.InvaliAugument(g,"token_expired")
		return
	}
	
	g.Next()
}