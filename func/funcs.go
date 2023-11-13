package funcs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
	return rn
}

func GenerateJWT2(email string, vCode int) (string, error) {
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &struct {
		Email string `json:"email"`
		jwt.RegisteredClaims
	}{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var key interface{}
	if vCode == 0 {
		key = []byte(os.Getenv("JWTKEY"))
	} else {
		key = []byte(strconv.Itoa(vCode))
	}
	return token.SignedString(key)
}

func VerifyJWT(tokenString string, vCode int) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(strconv.Itoa(vCode)), nil
	})

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}
