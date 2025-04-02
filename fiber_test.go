package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

func BenchmarkFiber(b *testing.B) {
	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "pong",
		})
	})

	//? Adaptor konversi fasthttp ke net/http agar dapat dilakukan pemanggilan ServeHTTP
	router := adaptor.FiberApp(app) 

	req := httptest.NewRequest("GET", "/ping", http.NoBody)
	for i := 0; i < b.N; i++ {
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
	}
}
