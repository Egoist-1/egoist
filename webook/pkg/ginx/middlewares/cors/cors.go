package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func CorsHandle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cors.New(cors.Config{
			AllowOriginFunc: func(origin string) bool {
				if strings.HasPrefix(origin, "http://localhost") {
					return true
				}
				return false
			},
			AllowHeaders:  []string{"Content-Type", "Authorization"},
			ExposeHeaders: []string{"jwt-token"},
			MaxAge:        time.Hour * 24,
		})
	}
}
