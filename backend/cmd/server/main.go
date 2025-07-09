package main

import (
	"cadence/internal/api"
	"cadence/internal/config"
	"cadence/internal/service"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// load env / defaults
	cfg := config.Load()

	// DI: create service clients
	getSongClient := service.NewGetSongClient(cfg.GetSongAPIKey)

	// init Gin
	r := gin.Default()
	r.Use(cors.New(config.CORS(cfg))) // CORS middleware

	// register REST routes
	api.RegisterGetSongRoutes(r, getSongClient)

	// start
	log.Println("⇢  Listening on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
