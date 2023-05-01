package middleWare

import (
	// "fmt"
	"golang/pkg/helpers"
	"golang/pkg/base"
	"github.com/gin-gonic/gin"
)

func InitJwtMiddleWare()(middleWare *JwtMiddleWare){
	var JwtMiddleWare JwtMiddleWare
 	return &JwtMiddleWare
}

type JwtMiddleWare struct {
	base.MiddleWare
}

func (m *JwtMiddleWare) ConfirmToken(g *gin.Context){
	jwtToken := g.Request.Header["Bearer-Token"][0]

	var token helpers.JwtToken

	val,err := token.IsValidJwt(jwtToken)
	if err!=nil{
		m.SystemError(g,err.Error())
		return
	}
	if !val {
		m.InvaliAugument(g,"invalid_token")
		return
	}

	val,err = token.IsJwtInTime(jwtToken)
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