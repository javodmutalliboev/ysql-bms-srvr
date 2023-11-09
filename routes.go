package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"ysql-bms/service"
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

	err := router.Run("localhost:3000")
	if err != nil {
		panic(err)
	}
}
