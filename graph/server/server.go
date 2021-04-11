package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/k88t76/GraphQL-gRPC-demo/article/client"
	"github.com/k88t76/GraphQL-gRPC-demo/graph"
	"github.com/k88t76/GraphQL-gRPC-demo/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// articleClientを生成
	articleClient, err := client.NewClient("localhost:50051")
	if err != nil {
		articleClient.Close()
		log.Fatalf("Failed to create article client: %v\n", err)
	}

	// GraphQLサーバーにResolverを登録
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					ArticleClient: articleClient,
				}}))

	// GraphQL playgroundのエンドポイント
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	//　実装したqueryが実行可能なサーバーのエンドポイント
	http.Handle("/query", srv)

	// GraphQLサーバーを起動
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
