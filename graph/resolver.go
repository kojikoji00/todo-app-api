package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate


import "github.com/kojikoji00/todo-app-api/graph/model"

type Resolver struct {
	todos []*model.Todo
}
