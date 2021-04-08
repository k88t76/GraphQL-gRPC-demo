package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	article "github.com/k88t76/GraphQL-gRPC-demo/article/client"
	"github.com/k88t76/GraphQL-gRPC-demo/graph"
)

const defaultPort = "8000"

func NewGraphQLServer(url string) (*graph.Resolver, error) {
	// articleサービスに接続
	articleClient, err := article.NewClient(url)
	if err != nil {
		articleClient.Close()
		return nil, err
	}
	return &graph.Resolver{
		Server: &graph.Server{
			ArticleClient: articleClient,
		}}, nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	s, err := NewGraphQLServer("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/query", handler.GraphQL(s.ToExecutableShema()))
	http.Handle("/", handler.Playground("Article", "/query"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
