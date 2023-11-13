package service

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strconv"
	funcs "ysql-bms/func"
)

func VerifyCode() func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody struct {
			Code int `json:"code" binding:"required"`
		}
		if err := c.ShouldBind(&requestBody); err != nil {
			funcs.ErrorResponse(c, err)
			return
		}

		signupToken, err := c.Cookie("signupToken")
		if err != nil {
			funcs.ErrorResponse(c, err)
			return
		}

		if err := funcs.VerifyJWT(signupToken, requestBody.Code); err != nil {
			funcs.ErrorResponse(c, err)
			return
		}

		token, err := jwt.Parse(signupToken, func(token *jwt.Token) (interface{}, error) {
			return strconv.Itoa(requestBody.Code), nil
		})

		email := token.Claims.(jwt.MapClaims)["email"].(string)

		passwordToken, _ := funcs.GenerateJWT2(email, 0)
		https, _ := strconv.ParseBool(os.Getenv("HTTPS"))
		c.SetCookie("signupToken", "", -100, "/", os.Getenv("DOMAIN"), https, true)
		c.SetCookie("passwordToken", passwordToken, 3600, "/", os.Getenv("DOMAIN"), https, true)
		c.JSON(200, gin.H{"verification code": "success"})
	}
}
