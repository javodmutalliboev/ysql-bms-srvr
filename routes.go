package main

import (
	"github.com/gin-gonic/gin"
	"ysql-bms/service"
)

func routes() {

	router := gin.Default()

	router.POST("/login", service.Login())

	err := router.Run("localhost:3000")
	if err != nil {
		panic(err)
	}
}
