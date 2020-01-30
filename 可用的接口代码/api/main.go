package main

import (
	_ "api/database"
	"api/router"
	orm "api/database"
)

func main() {
	defer orm.Eloquent.Close()
	router := router.InitRouter()
	router.Run(":8000")
}