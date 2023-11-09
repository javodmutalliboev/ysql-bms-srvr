package funcs

import (
	"fmt"
	"os"
	"strconv"
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
