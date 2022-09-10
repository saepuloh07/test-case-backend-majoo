package main

import (
	"fmt"
	"log"
	"net/http"
	"test-case-backend-majoo/handlers"
	"test-case-backend-majoo/middlewares"

	"test-case-backend-majoo/database"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Success connect to mysql")

	e := echo.New()

	// Validator
	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Use(middleware.Logger())  // Logger
	e.Use(middleware.Recover()) // Recover

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/login", handlers.Login)

	//
	r := e.Group("/report")
	r.Use(middleware.JWTWithConfig(middlewares.Config))
	r.GET("/merchants", handlers.GetReportByMerchants)
	r.GET("/outlets", handlers.GetReportByOutlets)

	e.Logger.Fatal(e.Start(":8080"))
}
