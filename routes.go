package main

import (
	"os"
	"ysql-bms/service"
	AdministratorService "ysql-bms/service/administrator"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func routes() {

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("CLIENTORIGIN")}
	config.AllowHeaders = []string{"Access-Control-Allow-Headers", "Authorization", "Content-Type"}
	config.AllowMethods = []string{"PATCH", "DELETE"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	router.POST("/login", service.Login())
	router.POST("/sendEmail", service.SendEmail())
	router.GET("/euList", service.GetExistingEmailList())
	router.POST("/verifyCode", service.VerifyCode())
	router.POST("/submitPassword", service.SubmitPassword())
	router.POST("/forgotPassword", service.ForgotPassword())

	administrator := router.Group("/administrator")
	{
		administrator.GET("/users", service.AuthAdmin(), AdministratorService.GetUsers())
	}

	router.GET("/logout", service.AuthUser(), service.Logout())

	err := router.Run("localhost:3000")
	if err != nil {
		panic(err)
	}
}
