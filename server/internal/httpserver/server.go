package httpserver

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func New(logger *slog.Logger) http.Handler {
	router := gin.New()
	router.Use(requestLogger(logger), gin.Recovery())

	router.GET("/health", health)

	return router
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
