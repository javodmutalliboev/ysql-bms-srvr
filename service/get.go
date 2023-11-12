package service

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	funcs "ysql-bms/func"
	_type "ysql-bms/type"
)

func GetExistingEmailList() func(c *gin.Context) {
	return func(c *gin.Context) {
		psqlInfo, err := funcs.PsqlInfo()

		if err != nil {
			funcs.ErrorResponse(c, err)
			return
		}

		euList, err := GetExistingEmailListDB(psqlInfo)
		if err != nil {
			funcs.ErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, euList)
	}
}

func GetExistingEmailListDB(psqlInfo string) (_type.EuList, error) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	sqlStatement := `SELECT email FROM public.user`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var euList _type.EuList
	for rows.Next() {
		var email string
		err = rows.Scan(&email)
		if err != nil {
			return nil, err
		}
		euList = append(euList, email)
	}
	return euList, nil
}
