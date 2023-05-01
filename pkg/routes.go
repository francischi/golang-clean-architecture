package pkg

import (
	"github.com/gin-gonic/gin"
	"golang/pkg/memberModule"
)

func SetRouter(g *gin.Engine) {

	const baseGroup string = "/api"
	memberModule.SetRoute(g ,baseGroup)
}