package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/sunilkkhadka/Forum/config"
	"github.com/sunilkkhadka/Forum/internal/routes"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.SetupRouter(r)

	srv := &http.Server{
		Addr:         cfg.Server.Addr,
		Handler:      r,
		WriteTimeout: cfg.Server.WriteTimeout,
		ReadTimeout:  cfg.Server.ReadTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	srv.ListenAndServe()

}
