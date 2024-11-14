package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, db *sql.DB) {
	api := router.Group("/api")
	RegisterV1Routes(api, db)
}
