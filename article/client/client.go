package main

import (
	"context"
	"fmt"
	"log"

	"github.com/k88t76/GraphQL-gRPC-demo/article/pb"
	"google.golang.org/grpc"
)

type Article struct {
	ID      int64  `json:"id"`
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Client struct {
	conn    *grpc.ClientConn
	service pb.ArticleServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := pb.NewArticleServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func main() {
	c, _ := NewClient("localhost:8080")
	/*input := &pb.CreateInput{
		Author:  "gopher",
		Title:   "gRPC-3",
		Content: "gRPC is so nice!",
	}
	res, err := c.service.CreateArticle(context.Background(), &pb.CreateArticleRequest{CreateInput: input})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("CreateArticle Response: %v", res)
	*/
	var id int64 = 1
	res, err := c.service.ReadArticle(context.Background(), &pb.ReadArticleRequest{Id: id})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ReadArticle Response: %v\n", res)
}

func (c *Client) CreateArticle(ctx context.Context, input *pb.CreateInput) (*Article, error) {
	res, err := c.service.CreateArticle(
		ctx,
		&pb.CreateArticleRequest{CreateInput: input},
	)
	if err != nil {
		return nil, err
	}
	return &Article{
		ID:      res.Article.Id,
		Author:  res.Article.Author,
		Title:   res.Article.Title,
		Content: res.Article.Content,
	}, nil
}

func (c *Client) ReadArticle(ctx context.Context, id int64) (*Article, error) {
	res, err := c.service.ReadArticle(ctx, &pb.ReadArticleRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &Article{
		ID:      res.Article.Id,
		Author:  res.Article.Author,
		Title:   res.Article.Title,
		Content: res.Article.Content,
	}, nil
}
