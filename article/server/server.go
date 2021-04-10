package main

import (
	"log"
	"net"

	"github.com/k88t76/GraphQL-gRPC-demo/article/pb"
	"github.com/k88t76/GraphQL-gRPC-demo/article/repository"
	"github.com/k88t76/GraphQL-gRPC-demo/article/service"
	"google.golang.org/grpc"
)

func main() {

	// articleサーバーに接続
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer lis.Close()

	// RepositoryとServiceを作成
	repository, err := repository.NewsqliteRepo()
	if err != nil {
		log.Fatalf("Failed to create new sqlite repository: %v\n", err)
	}

	service := service.NewService(repository)
	//サーバーにarticleサービスを登録
	server := grpc.NewServer()
	pb.RegisterArticleServiceServer(server, service)

	//articleサーバーを起動
	log.Println("Listening on port 8080...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}

}
