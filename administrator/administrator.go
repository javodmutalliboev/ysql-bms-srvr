package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"reflect"
	"strconv"
)

func practice() {
	fmt.Println("working")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
		return
	}

	//port, _ := strconv.Atoi(os.Getenv("PORT"))
	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	os.Getenv("HOST"), port, os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))

	//createAdminUser(psqlInfo)
	password := os.Getenv("ADMINPASSWORD")

	hash, _ := HashPassword(password)

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, "$2a$14$xS8Q71gH.jNbR5A/0vvUreip0z0gxbuFdwU8dTH4p1kahlPRWmBy.")
	fmt.Println("Match:   ", match)

	//editAdminUser(psqlInfo, hash)
	https, _ := strconv.ParseBool(os.Getenv("HTTPS"))
	log.Println(reflect.TypeOf(https))
}

func HashPassword(password string) (string, error) {
	log.Println(password)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func editAdminUser(psqlInfo string, hash string) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sqlStatement := `
		UPDATE public.user
		SET password = $2
		WHERE email = $1
	`
	_, err2 := db.Exec(sqlStatement, "javodmutalliboev@gmail.com", hash)
	if err2 != nil {
		panic(err2)
	}
	log.Printf("password for the admin user with email %v has been set", "javodmutalliboev@gmail.com")
}

func createAdminUser(psqlInfo string) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO public.user (email, first_name, last_name, role)
		VALUES ($1, $2, $3, $4)
		RETURNING email`
	email := "javodmutalliboev@gmail.com"
	err = db.QueryRow(sqlStatement, email, "Javod", "Mutalliboev", "administrator").Scan(&email)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", email)
}
