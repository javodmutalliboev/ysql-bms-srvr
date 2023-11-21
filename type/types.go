package _type

import (
	"database/sql"
	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

type AuthenticationResult struct {
	Email     sql.NullString `json:"email"`
	FirstName sql.NullString `json:"first_name"`
	LastName  sql.NullString `json:"last_name"`
	Role      sql.NullString `json:"role"`
}

type EuList []string
