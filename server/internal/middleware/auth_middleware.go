package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunilkkhadka/Forum/internal/utils"
	"github.com/sunilkkhadka/Forum/internal/utils/auth"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, err := ctx.Cookie("access_token")
		if err != nil {
			utils.SendErrorResponse(ctx, utils.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "Token Not Found",
			})
			return
		}

		claims, err := auth.ValidateToken(accessToken, []byte(auth.JwtConf.JwtAccessTokenSecret))
		if err != nil {
			utils.SendErrorResponse(ctx, utils.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "Invalid Token",
			})
			return
		}

		ctx.Set("user_id", claims["user_id"])

		ctx.Next()
	}
}
