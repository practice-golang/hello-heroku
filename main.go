package main // import "hello-heroku"

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func showAllEnv(c echo.Context) error {
	envs := os.Environ()
	return c.JSON(http.StatusOK, envs)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		// log.Fatal("$PORT must be set")
		port = "1323"
	}

	e := echo.New()
	e.Use(middleware.Recover())
	// e.Use(middleware.Logger())
	e.HideBanner = true

	e.GET("/", hello)
	e.GET("/health", healthCheck)
	e.GET("/datetime", showDateTime)
	e.GET("/env", showAllEnv)

	e.Logger.Fatal(e.Start(":" + port))
}
