package httpserver

import (
	"bytes"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRequestLogger(t *testing.T) {
	gin.SetMode(gin.TestMode)

	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, nil))

	router := gin.New()
	router.Use(requestLogger(logger))
	router.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusCreated)
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	router.ServeHTTP(httptest.NewRecorder(), req)

	out := buf.String()
	assert.Contains(t, out, "msg=request")
	assert.Contains(t, out, "method=GET")
	assert.Contains(t, out, "path=/test")
	assert.Contains(t, out, "status=201")
}
