// package main

// import (
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/99designs/gqlgen/graphql/handler"
// 	"github.com/99designs/gqlgen/graphql/playground"
// 	"github.com/kojikoji00/todo-app-api/graph"
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// )

// const defaultPort = "8080"

// func main() {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = defaultPort
// 	}

// 	e := echo.New()

// 	// GraphQL handler
// 	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
// 	e.POST("/query", echo.WrapHandler(srv))

// 	// GraphQL playground
// 	e.GET("/", echo.WrapHandler(playground.Handler("GraphQL playground", "/query")))

// 	// Echo routes
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	e.GET("/hello", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Hello, Echo!")
// 	})

// 	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
// 	log.Fatal(e.Start(":" + port))
// }