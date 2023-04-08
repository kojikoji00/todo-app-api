// package main

// import (
// 	"net/http"

// 	_ "github.com/99designs/gqlgen"
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// )

// func main() {
// 	e := echo.New()

// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	e.GET("/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Hello, Echo!")
// 	})

// 	e.Start(":8080")
// }
