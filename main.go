package main

import (
	"internship-stikom/config"
	"internship-stikom/routes"
)

func main() {
	config.ConnectDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}
