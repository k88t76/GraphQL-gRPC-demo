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

func main() {
	c, _ := NewClient("localhost:8080")
	/*
		input := &pb.CreateInput{
			Author:  "gopher3",
			Title:   "poyo",
			Content: "gRPC is poyopoyo",
		}
		res, err := c.service.CreateArticle(context.Background(), &pb.CreateArticleRequest{CreateInput: input})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("CreateArticle Response: %v", res)
	*/

	/*
		var id int64 = 1
		res, err := c.service.ReadArticle(context.Background(), &pb.ReadArticleRequest{Id: id})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ReadArticle Response: %v\n", res)
	*/
	/*
		article := &pb.Article{
			Id:      2,
			Author:  "gopher2",
			Title:   "GraphQL",
			Content: "GraphQL is very very smart!",
		}
		res, err := c.service.UpdateArticle(context.Background(), &pb.UpdateArticleRequest{Article: article})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("UpdateArticle Response: %v\n", res)
	*/
	/*
		var id int64 = 4
		res, err := c.service.DeleteArticle(context.Background(), &pb.DeleteArticleRequest{Id: id})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Deleted Article id = %v\n", res.Id)
	*/
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

func (c *Client) UpdateArticle(ctx context.Context, article *pb.Article) (*Article, error) {
	res, err := c.service.UpdateArticle(ctx, &pb.UpdateArticleRequest{Article: article})
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

func (c *Client) DeleteArticle(ctx context.Context, id int64) (int64, error) {
	res, err := c.service.DeleteArticle(ctx, &pb.DeleteArticleRequest{Id: id})
	if err != nil {
		return 0, err
	}
	return res.Id, nil
}
