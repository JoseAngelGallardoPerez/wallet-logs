package main

import (
	"log"

	"github.com/Confialink/wallet-logs/internal/logs/config"
	"github.com/Confialink/wallet-logs/internal/logs/protobuf_server"
	"github.com/Confialink/wallet-logs/internal/logs/routes"
	"github.com/Confialink/wallet-pkg-env_mods"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.GetConfig()
	ginMode := env_mods.GetMode(config.Env)
	gin.SetMode(ginMode)
	router := routes.GetRouter()

	protobuf_server.StartProtobufServer()

	log.Printf("Starting API on port: %s", config.Port)
	router.Run(":" + config.Port)
}
