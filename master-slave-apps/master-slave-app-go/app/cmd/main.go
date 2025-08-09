package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"apps/repositories"
	"apps/routes"
	"apps/config"
	"apps/handlers"
)

func main() {
	// Muat konfigurasi dari .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Gagal memuat file .env")
	}

	// Inisialisasi konfigurasi database
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal("Gagal inisialisasi konfigurasi: ", err)
	}

	// Inisialisasi repository
	repo, err := repositories.NewItemRepository(config)
	if err != nil {
		log.Fatal("Gagal inisialisasi repository: ", err)
	}



	// Inisialisasi router
	r := gin.Default()
	routes.SetupRoutes(r, handlers.NewItemHandler(repo))

	// Jalankan server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}