package main // import "hello-heroku"

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

var (
	//go:embed static
	content embed.FS
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world!")
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

func dbHealth(c echo.Context, db *sql.DB) error {
	var msg string

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS ticks (tick timestamp)"); err != nil {
		msg = fmt.Sprintf("Error creating database table: %q", err)
		return c.String(http.StatusInternalServerError, msg)
	}

	if _, err := db.Exec("INSERT INTO ticks VALUES (now())"); err != nil {
		msg = fmt.Sprintf("Error incrementing tick: %q", err)
		return c.String(http.StatusInternalServerError, msg)
	}

	rows, err := db.Query("SELECT tick FROM ticks")
	if err != nil {
		msg = fmt.Sprintf("Error reading ticks: %q", err)
		return c.String(http.StatusInternalServerError, msg)
	}

	defer rows.Close()

	msg = ""
	for rows.Next() {
		var tick time.Time
		if err := rows.Scan(&tick); err != nil {
			msg = fmt.Sprintf("Error scanning ticks: %q", err)
			return c.String(http.StatusInternalServerError, msg)
		}

		msg += fmt.Sprintf("Read from DB: %s\n", tick.String())
	}

	return c.String(http.StatusOK, msg)
}

func tableClear(c echo.Context, db *sql.DB) error {
	var msg string
	if _, err := db.Exec("TRUNCATE ticks"); err != nil {
		msg = fmt.Sprintf("Error incrementing tick: %q", err)
		return c.String(http.StatusInternalServerError, msg)
	}

	return c.String(http.StatusOK, "done")
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		// log.Fatal("$PORT must be set")
		port = "1323"
	}

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.HideBanner = true

	contentHandler := echo.WrapHandler(http.FileServer(http.FS(content)))
	contentRewrite := middleware.Rewrite(map[string]string{"/*": "/static/$1"})

	e.GET("/*", contentHandler, contentRewrite)
	e.GET("/hello", hello)
	e.GET("/health", healthCheck)
	e.GET("/datetime", showDateTime)
	e.GET("/env", showAllEnv)

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	dbHandler := func(c echo.Context) error {
		return dbHealth(c, db)
	}
	tableClear := func(c echo.Context) error {
		return tableClear(c, db)
	}
	e.GET("/db", dbHandler)
	e.GET("/table-clear", tableClear)

	e.Logger.Fatal(e.Start(":" + port))
}
