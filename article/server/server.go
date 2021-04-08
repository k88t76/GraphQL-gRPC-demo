package server

import (
	"context"
	"database/sql"
	"log"
	"net"

	"github.com/k88t76/GraphQL-gRPC-demo/article/pb"
	"google.golang.org/grpc"
)

type server struct {
}

type articleInput struct {
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var db *sql.DB

func main() {
	// sqliteに接続
	var err error
	db, err = sql.Open("sqlite3", "./blog/blog.sql")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// articlesテーブルを作成
	cmd := `CREATE TABLE IF NOT EXISTS articles(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		author STRING,
		content STRING,
		title STRING)`

	_, err = db.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// articleサーバーに接続
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer lis.Close()
	s := grpc.NewServer()

	//サーバーにarticleサービスを登録
	pb.RegisterArticleServiceServer(s, &server{})

	//articleサーバーを起動
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}

}

func (*server) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
	input := req.GetCreateInput()

	cmd := "INSERT INTO articles(author, title, content) VALUES (?, ?, ?)"
	r, err := db.Exec(cmd, input.Author, input.Title, input.Content)
	if err != nil {
		return nil, err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &pb.CreateArticleResponse{
		Article: &pb.Article{
			Id:      id,
			Author:  input.Author,
			Title:   input.Title,
			Content: input.Content,
		},
	}, nil
}
