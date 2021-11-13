package graph

//go:generate go run github.com/99designs/gqlgen

import "github.com/ryokubozono/go-gqlgen/graph/model"

type Resolver struct {
	todos []*model.Todo
}
