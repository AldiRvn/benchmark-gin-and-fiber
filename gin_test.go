package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func BenchmarkGin(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/ping", http.NoBody)
	for i := 0; i < b.N; i++ {
		router.ServeHTTP(recorder, request)
	}
}
