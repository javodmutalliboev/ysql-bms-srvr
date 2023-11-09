package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	routes()
}
