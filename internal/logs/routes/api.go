package routes

import (
	"net/http"

	"github.com/Confialink/wallet-logs/internal/logs/http/middlewares"
	"github.com/Confialink/wallet-logs/internal/version"

	"github.com/Confialink/wallet-logs/internal/authentication"
	"github.com/Confialink/wallet-pkg-errors"
	"github.com/Confialink/wallet-pkg-service_names"
	"github.com/gin-gonic/gin"

	"github.com/Confialink/wallet-logs/internal/logs/config/logs"
	"github.com/Confialink/wallet-logs/internal/logs/http/app/common"
	"github.com/Confialink/wallet-logs/internal/logs/http/app/cors"
	"github.com/Confialink/wallet-logs/internal/logs/http/handlers"
)

var r *gin.Engine

func initRoutes() {
	r = gin.New()

	// Middleware

	mwAuth := authentication.Middleware(logs.Logger.New("Middleware", "authentication"))

	// Handlers

	corsHndlr := cors.NewCorsHandler()
	notFoundHndlr := common.NewNotFoundHandler()
	logsHndlr := handlers.NewLogsHandler()
	csvHndlr := handlers.NewCsvLogsHandler()

	// Routes

	r.GET("/logs/health-check", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	r.GET("/logs/build", func(c *gin.Context) {
		c.JSON(http.StatusOK, version.BuildInfo)
	})

	apiGroup := r.Group(service_names.Logs.Internal, cors.CorsMiddleware())
	apiGroup.Use(
		gin.Recovery(),
		gin.Logger(),
		errors.ErrorHandler(logs.Logger.New("Middleware", "errors")),
	)

	privateGroup := apiGroup.Group("/private", mwAuth, middlewares.CanViewSystemLogs())
	{
		v1Group := privateGroup.Group("/v1")
		{
			adminGroup := v1Group.Group("/admin")
			adminGroup.GET("system-logs", logsHndlr.List)
			adminGroup.GET("system-logs/:id", logsHndlr.Get)

			adminExportGroup := adminGroup.Group("/export")
			{
				adminExportGroup.GET("transactions-log", csvHndlr.DownloadTransactions)
				adminExportGroup.GET("information-log", csvHndlr.DownloadInformation)
			}
		}
	}

	// Handle OPTIONS request
	r.OPTIONS("/*cors", corsHndlr.OptionsHandler, cors.CorsMiddleware())

	r.NoRoute(notFoundHndlr.NotFoundHandler)
}

func GetRouter() *gin.Engine {
	if nil == r {
		initRoutes()
	}
	return r
}
