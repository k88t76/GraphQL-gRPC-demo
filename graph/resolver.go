package graph

import "github.com/k88t76/GraphQL-gRPC-demo/article/client"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ArticleClient *client.Client
}
