package main

import (
	"assignment2/database"
	"assignment2/routers"
)

func main() {
	database.InitializeDB()

	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
