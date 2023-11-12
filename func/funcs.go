package funcs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"time"
)

func PsqlInfo() (string, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		return "", err
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"), port, os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))

	return psqlInfo, nil
}

func ErrorResponse(c *gin.Context, err error) {
	log.Println(err)
	c.String(http.StatusInternalServerError, err.Error())
}

func SendVerificationCode(to string, message string) error {
	auth := smtp.PlainAuth("", os.Getenv("SMTPUSER"), os.Getenv("SMTPPASSWORD"), os.Getenv("SMTPSERVER"))
	err := smtp.SendMail(os.Getenv("SMTPSERVER")+":"+os.Getenv("SMTPPORT"), auth, os.Getenv("SMTPUSER"), []string{to}, []byte(message))
	if err != nil {
		return err
	}
	return nil
}

func GenerateRandomNumber() int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	minRn := 100000
	maxRn := 999999
	rn := rand.Intn(maxRn-minRn+1) + minRn
	RandomNumbers = append(RandomNumbers, rn)
	return rn
}

var RandomNumbers []int
