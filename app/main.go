package main

import (
	"goproj2/api"
	"goproj2/db"
)

func main() {
	// Connect to database
	db.ConnectDatabase()

	r := api.SetupRouter()

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run()
}