package graph

import (
	"github.com/99designs/gqlgen/graphql"
	article "github.com/k88t76/GraphQL-gRPC-demo/article/client"
	"github.com/k88t76/GraphQL-gRPC-demo/graph/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Server struct {
	ArticleClient *article.Client
}

type Resolver struct {
	Server *Server
}

// スキーマを実行できるようにするメソッド
func (r *Resolver) ToExecutableShema() graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: r,
	})
}
