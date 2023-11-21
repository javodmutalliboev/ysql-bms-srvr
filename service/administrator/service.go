package AdministratorService

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	funcs "ysql-bms/func"
)

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		psqlInfo, err := funcs.PsqlInfo()
		if err != nil {
			funcs.ErrorResponse(c, err)
			c.Abort()
			return
		}

		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			funcs.ErrorResponse(c, err)
			c.Abort()
			return
		}

		defer db.Close()

		rows, err := db.Query("SELECT * FROM public.user")
		if err != nil {
			funcs.ErrorResponse(c, err)
			c.Abort()
			return
		}

		var users []map[string]string
		for rows.Next() {
			var email, first_name, last_name, role, password sql.NullString
			err = rows.Scan(&email, &first_name, &last_name, &role, &password)
			if err != nil {
				funcs.ErrorResponse(c, err)
				c.Abort()
				return
			}
			user := map[string]string{
				"email":      email.String,
				"first_name": first_name.String,
				"last_name":  last_name.String,
				"role":       role.String,
			}
			users = append(users, user)
		}

		c.JSON(200, gin.H{"users": users})
	}
}
