package main

import (
	"context"

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

func (c *Client) CreateArticle(ctx context.Context, input *pb.CreateInput) (*Article, error) {
	r, err := c.service.CreateArticle(
		ctx,
		&pb.CreateArticleRequest{CreateInput: input},
	)
	if err != nil {
		return nil, err
	}
	return &Article{
		ID:      r.Article.Id,
		Author:  r.Article.Author,
		Title:   r.Article.Title,
		Content: r.Article.Content,
	}, nil
}
