package router

import (
	"net/http"
	"seth16888/api/gw/internal/server/handler"
	"seth16888/api/gw/internal/server/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	// gin middleware
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
)

func NewRouter(log *zap.Logger) http.Handler {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(middleware.LoggingMiddleware(log))
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(requestid.New())

	registerRoutes(r)

	return r
}

func registerRoutes(r *gin.Engine) {
	r.GET("/ping", handler.NewHealthHandler().Ping)
	r.GET("/health", handler.NewHealthHandler().Health)
}
