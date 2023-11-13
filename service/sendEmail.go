package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
	funcs "ysql-bms/func"
)

func SendEmail() func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody struct {
			Email string `json:"email" binding:"required,email"`
		}

		if err := c.ShouldBind(&requestBody); err != nil {
			println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
			return
		}

		psqlInfo, err := funcs.PsqlInfo()
		if err != nil {
			funcs.ErrorResponse(c, err)
			return
		}
		l, err := GetExistingEmailListDB(psqlInfo)
		if err != nil {
			funcs.ErrorResponse(c, err)
			return
		}

		for _, v := range l {
			if v == requestBody.Email {
				log.Printf("email already exists: %s\n", requestBody.Email)
				c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
				return
			}
		}

		vCode := funcs.GenerateRandomNumber()
		message := "Subject: [Book Management System] Email Verification Code\n\n" + "Your verification code is " + strconv.Itoa(vCode) + "."
		err = funcs.SendVerificationCode(requestBody.Email, message)

		if err != nil {
			funcs.ErrorResponse(c, err)
			return
		}

		token, _ := funcs.GenerateJWT2(requestBody.Email, vCode)

		https, _ := strconv.ParseBool(os.Getenv("HTTPS"))
		c.SetCookie("signupToken", token, 3600, "/", os.Getenv("DOMAIN"), https, true)

		c.JSON(http.StatusOK, gin.H{"verification code send": "success"})
	}
}
