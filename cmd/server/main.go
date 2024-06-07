package main

import (
	"log"
	"net/http"
	"os"
	"ozon-posts/graph"
	"ozon-posts/internal/repo"
	"ozon-posts/internal/repo/db"
	"ozon-posts/internal/repo/memory"
	"ozon-posts/internal/service"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {

	storage := os.Getenv("STORAGE_TYPE")
	if storage != "db" && storage != "memory" {
		log.Fatal(
			"unknown storage type",
		)
	}

	var postRepository repo.PostRepo
	var commentRepository repo.CommentRepo

	if storage == "db" {

		c, cErr := db.Connection()
		if cErr != nil {
			log.Fatal(
				cErr,
				"creation of the db connection failed",
			)
		}
		defer c.Close()

		postRepository = db.NewPostDbRepository(c)
		commentRepository = db.NewCommentDbRepository(c)
	} else {
		postRepository = memory.NewPostRepository()
		commentRepository = memory.NewCommentRepository()
	}

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: graph.NewResolver(
					service.NewPostService(postRepository),
					service.NewCommentService(commentRepository),
				),
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL  playground", "/query"))
	http.Handle("/query", srv)

	log.Print("connect to http://localhost:8080/ for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
