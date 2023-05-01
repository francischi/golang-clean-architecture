package base

import (
	// "fmt"
	// "errors"
	"github.com/gin-gonic/gin"
	"golang/pkg/except"
)

type Controller struct {
}

func (c *Controller) SuccessRes(g *gin.Context ,result interface{}) {
	g.JSON(200,gin.H{
		"success":true,
		"error_msg":nil,
		"result":result,
	})
}

func (c *Controller)HandleError(g *gin.Context , err error){
	statusCode := except.Error2StatusCode(err)
	g.JSON(statusCode,gin.H{
		"success":false,
		"error_msg":err.Error(),
		"result":nil,
	})
}