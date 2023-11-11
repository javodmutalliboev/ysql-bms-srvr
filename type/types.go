package _type

import "github.com/golang-jwt/jwt/v4"

type User struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Claims struct {
	Email string `json:"emil"`
	jwt.RegisteredClaims
}

type AuthenticationResult struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
}

type EuList []string
