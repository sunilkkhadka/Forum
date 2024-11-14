package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/sunilkkhadka/Forum/config"
	"github.com/sunilkkhadka/Forum/internal/database"
	"github.com/sunilkkhadka/Forum/internal/routes"
)

func main() {
	// Load Configurations
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	log.Println("Configurations Loaded Successfully!!")

	// Connect to Database
	db, err := database.InitDB(cfg.Database)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	log.Println("Database Connected Successfully!!")
	defer db.Close()

	// Initiate Router
	r := gin.Default()

	// Allow Cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Setup Router
	routes.SetupRouter(r, db)

	// Configure Server
	srv := &http.Server{
		Addr:         cfg.Server.Addr,
		Handler:      r,
		WriteTimeout: cfg.Server.WriteTimeout,
		ReadTimeout:  cfg.Server.ReadTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	// Create Server
	srv.ListenAndServe()
}
