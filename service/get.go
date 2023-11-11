package service

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	funcs "ysql-bms/func"
	_type "ysql-bms/type"
)

func GetExistingEmailList() func(c *gin.Context) {
	return func(c *gin.Context) {
		psqlInfo, err := funcs.PsqlInfo()

		if err != nil {
			log.Printf("An error occurrred: %v.", err)
			c.String(http.StatusInternalServerError, "An error occurred.")
			return
		}

		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Printf("An error occurrred: %v.", err)
			c.String(http.StatusInternalServerError, "An error occurred.")
			return
		}
		defer db.Close()
		sqlStatement := `SELECT email FROM public.user`
		rows, err := db.Query(sqlStatement)
		if err != nil {
			log.Printf("An error occurrred: %v.", err)
			c.String(http.StatusInternalServerError, "An error occurred.")
			return
		}
		defer rows.Close()
		var euList _type.EuList
		for rows.Next() {
			var email string
			err = rows.Scan(&email)
			if err != nil {
				log.Printf("An error occurrred: %v.", err)
				c.String(http.StatusInternalServerError, "An error occurred.")
				return
			}
			euList = append(euList, email)
		}
		c.JSON(http.StatusOK, euList)
	}
}
