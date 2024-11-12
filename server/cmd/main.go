package main

import (
	"log"
	"net/http"

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
