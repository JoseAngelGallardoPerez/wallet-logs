package cors

import (
	"github.com/Confialink/wallet-logs/internal/logs/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	cfg := config.GetConfig()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowMethods = cfg.Cors.Methods

	for _, origin := range cfg.Cors.Origins {
		if origin == "*" {
			corsConfig.AllowAllOrigins = true
		}
	}
	if !corsConfig.AllowAllOrigins {
		corsConfig.AllowOrigins = cfg.Cors.Origins
	}
	corsConfig.AllowHeaders = cfg.Cors.Headers

	return cors.New(corsConfig)
}
