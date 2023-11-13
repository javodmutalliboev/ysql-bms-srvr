package service

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
	"regexp"
	"strconv"
	"ysql-bms/administrator"
	funcs "ysql-bms/func"
)

func SubmitPassword() func(c *gin.Context) {
	return func(c *gin.Context) {
		token, err := c.Cookie("passwordToken")
		if err != nil {
			funcs.ErrorResponse(c, err)
			return
		}

		var requestBody struct {
			Password string `json:"password" binding:"required,min=8,max=20"`
		}

		if err := c.ShouldBind(&requestBody); err != nil {
			funcs.ErrorResponse(c, err)
			return
		}

		_, err = regexp.MatchString(`^\w{8,}$`, requestBody.Password)
		if err != nil {
			funcs.ErrorResponse(c, err)
			return
		}

		psqlInfo, err := funcs.PsqlInfo()
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			funcs.ErrorResponse(c, err)
			return
		}
		defer db.Close()
		sqlStatement := `
			INSERT INTO public.user (email, role, password)
			VALUES ($1, $2, $3)
		`
		tokenParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWTKEY")), nil
		})
		if err != nil {
			log.Println("here")
			funcs.ErrorResponse(c, err)
			return
		}
		email := tokenParsed.Claims.(jwt.MapClaims)["email"].(string)
		role := "subscriber"
		password, _ := administrator.HashPassword(requestBody.Password)
		_, err = db.Exec(sqlStatement, email, role, password)
		if err != nil {
			funcs.ErrorResponse(c, err)
			return
		}
		https, _ := strconv.ParseBool(os.Getenv("HTTPS"))
		c.SetCookie("passwordToken", "", -100, "/", os.Getenv("DOMAIN"), https, true)
		c.JSON(200, gin.H{"signup": "success"})
	}
}
