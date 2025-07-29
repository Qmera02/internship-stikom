package main

import (
	"internship-stikom/config"
	"internship-stikom/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// HARUS load .env paling awal
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal load .env:", err)
	}

	log.Println("JWT_SECRET =", os.Getenv("JWT_SECRET")) // ✅ Tambahkan ini

	config.ConnectDB()

	r := routes.SetupRouter()
	r.Run(":8080")
}
