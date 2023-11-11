package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Signup() func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody struct {
			Email string `json:"email" binding:"required,email"`
		}

		if err := c.ShouldBind(&requestBody); err != nil {
			println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
			return
		}
	}
}
