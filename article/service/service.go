package service

import (
	"context"

	"github.com/k88t76/GraphQL-gRPC-demo/article/pb"
	"github.com/k88t76/GraphQL-gRPC-demo/article/repository"
)

type Service interface {
	CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error)
	ReadArticle(ctx context.Context, req *pb.ReadArticleRequest) (*pb.ReadArticleResponse, error)
	UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleResponse, error)
	DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleResponse, error)
	ListArticle(req *pb.ListArticleRequest, stream pb.ArticleService_ListArticleServer) error
}

type articleService struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return &articleService{r}
}

func (s *articleService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
	input := req.GetArticleInput()
	id, err := s.repository.InsertArticle(ctx, input)
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

func (s *articleService) ReadArticle(ctx context.Context, req *pb.ReadArticleRequest) (*pb.ReadArticleResponse, error) {
	id := req.GetId()
	a, err := s.repository.SelectArticleByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &pb.ReadArticleResponse{
		Article: &pb.Article{
			Id:      a.Id,
			Author:  a.Author,
			Title:   a.Title,
			Content: a.Content,
		},
	}, nil

}

func (s *articleService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleResponse, error) {
	id := req.GetId()
	input := req.GetArticleInput()

	a, err := s.repository.UpdateArticle(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateArticleResponse{
		Article: &pb.Article{
			Id:      a.Id,
			Author:  a.Author,
			Title:   a.Title,
			Content: a.Content,
		},
	}, nil
}

func (s *articleService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleResponse, error) {
	id := req.GetId()
	if err := s.repository.DeleteArticle(ctx, id); err != nil {
		return nil, err
	}
	return &pb.DeleteArticleResponse{Id: id}, nil
}

func (s *articleService) ListArticle(req *pb.ListArticleRequest, stream pb.ArticleService_ListArticleServer) error {
	rows, err := s.repository.SelectAllArticles()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var a pb.Article
		err := rows.Scan(&a.Id, &a.Author, &a.Title, &a.Content)
		if err != nil {
			return err
		}
		stream.Send(&pb.ListArticleResponse{Article: &a})
	}
	return nil
}
