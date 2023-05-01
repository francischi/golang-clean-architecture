package main

import (
	// "fmt"
	"golang/pkg"
	"golang/pkg/helpers"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	if err,_ := helpers.InitMySql();err !=nil{
		panic(err)
	}

	if err, _ := helpers.InitRedisConn();err !=nil{
		panic(err)
	}

	if err := helpers.InitEnvSetting();err !=nil{
		panic(err)
	}

	router := gin.New()
	router.Use(gin.Logger())
	pkg.SetRouter(router)
	router.GET("/" , func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"success" : true, 
			"message" : "service is running", 
		})
	})

	router.Run(":8000")
}