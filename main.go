package main // import "hello-heroku"

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

func showDateTime(c echo.Context) error {
	now := time.Now().Format("2006-01-02 15:04:05")
	return c.String(http.StatusOK, now)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		// log.Fatal("$PORT must be set")
		port = "1323"
	}

	e := echo.New()
	e.HideBanner = true

	e.GET("/", hello)
	e.GET("/health", healthCheck)
	e.GET("/datetime", showDateTime)
	e.Logger.Fatal(e.Start(":" + port))
}
