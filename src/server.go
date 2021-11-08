package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
)

const dataSource = "localuser:localpass@tcp(127.0.0.1:3306)/localdb?charset=utf8&parseTime=True&loc=Local"
const defaultPort = "5050"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := gorm.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic(err)
	}
	defer func() {
		if db != nil {
			if err := db.Close(); err != nil {
				panic(err)
			}
		}
	}()
	db.LogMode(true)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query",
		handler.NewDefaultServer(
			generated.NewExecutableSchema(
				generated.Config{
					Resolvers: &graph.Resolver{
						DB: db,
					},
				},
			),
		),
	)
}
