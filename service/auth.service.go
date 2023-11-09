package service

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	funcs "ysql-bms/func"
	_type "ysql-bms/type"
)

var Tokens []string

func Login() func(c *gin.Context) {
	return func(c *gin.Context) {
		if user, err := c.Cookie("user"); err == nil {
			log.Printf("alredy logged in: user: %s\n", user)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
			return
		}

		if token, err := c.Cookie("token"); err == nil {
			log.Printf("alredy logged in: token: %s\n", token)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
			return
		}

		var requestBody _type.User

		if err := c.ShouldBind(&requestBody); err != nil {
			println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
			return
		}

		data, err := AuthenticateUser(_type.User{Email: requestBody.Email, Password: requestBody.Password})
		if err != nil {
			log.Println(err)
			c.String(http.StatusNotImplemented, err.Error())
			return
		}

		token, _ := GenerateJWT(data.Email)
		Tokens = append(Tokens, token)

		https, _ := strconv.ParseBool(os.Getenv("HTTPS"))
		user, err := json.Marshal(map[string]string{"email": data.Email, "first_name": data.FirstName, "last_name": data.LastName, "role": data.Role})
		c.SetCookie("token", token, 3600, "/", os.Getenv("DOMAIN"), https, true)
		c.SetCookie("user", string(user), 3600, "/", os.Getenv("DOMAIN"), https, false)
		c.JSON(http.StatusOK, gin.H{"login": "success"})
	}
}

func AuthenticateUser(user _type.User) (_type.AuthenticationResult, error) {
	psqlInfo, err := funcs.PsqlInfo()
	if err != nil {
		return _type.AuthenticationResult{}, err
	}
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return _type.AuthenticationResult{}, err
	}
	defer db.Close()
	sqlStatement := `SELECT * FROM public.user WHERE email = $1`
	row := db.QueryRow(sqlStatement, user.Email)
	var email, first_name, last_name, role, password string
	switch err := row.Scan(&email, &first_name, &last_name, &role, &password); err {
	case sql.ErrNoRows:
		return _type.AuthenticationResult{}, err
	case nil:
		authenticated := CheckPasswordHash(user.Password, password)
		if authenticated {
			return _type.AuthenticationResult{email, first_name, last_name, role}, nil
		} else {
			return _type.AuthenticationResult{}, errors.New("unauthenticated")
		}
	default:
		return _type.AuthenticationResult{}, err
	}
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &_type.Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWTKEY")))

}
