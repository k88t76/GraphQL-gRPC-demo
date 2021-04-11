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

type service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{r}
}

func (s *service) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
	// INSERTする記事のInput(Author, Title, Content)を取得
	input := req.GetArticleInput()

	// 記事をDBにINSERTし、INSERTした記事のIDを返す
	id, err := s.repository.InsertArticle(ctx, input)
	if err != nil {
		return nil, err
	}

	// INSERTした記事をレスポンスとして返す
	return &pb.CreateArticleResponse{
		Article: &pb.Article{
			Id:      id,
			Author:  input.Author,
			Title:   input.Title,
			Content: input.Content,
		},
	}, nil
}

func (s *service) ReadArticle(ctx context.Context, req *pb.ReadArticleRequest) (*pb.ReadArticleResponse, error) {
	// READする記事のIDを取得
	id := req.GetId()

	// DBから該当IDの記事を取得
	a, err := s.repository.SelectArticleByID(ctx, id)
	if err != nil {
		return nil, err
	}

	//　取得した記事をレスポンスとして返す
	return &pb.ReadArticleResponse{
		Article: &pb.Article{
			Id:      id,
			Author:  a.Author,
			Title:   a.Title,
			Content: a.Content,
		},
	}, nil

}

func (s *service) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleResponse, error) {
	// UPDATEする記事のIDを取得
	id := req.GetId()

	// UPDATEする記事の変更内容(Author, Title, Content)を取得
	input := req.GetArticleInput()

	//　該当IDの記事をUPDATE
	if err := s.repository.UpdateArticle(ctx, id, input); err != nil {
		return nil, err
	}

	// UPDATEした記事をレスポンスとして返す
	return &pb.UpdateArticleResponse{
		Article: &pb.Article{
			Id:      id,
			Author:  input.Author,
			Title:   input.Title,
			Content: input.Content,
		},
	}, nil
}

func (s *service) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleResponse, error) {
	// DELETEする記事のIDを取得
	id := req.GetId()

	// 該当IDの記事をDELETE
	if err := s.repository.DeleteArticle(ctx, id); err != nil {
		return nil, err
	}

	// DELETEした記事のIDをレスポンスとして返す
	return &pb.DeleteArticleResponse{Id: id}, nil
}

func (s *service) ListArticle(req *pb.ListArticleRequest, stream pb.ArticleService_ListArticleServer) error {
	// 記事を全取得
	rows, err := s.repository.SelectAllArticles()
	if err != nil {
		return err
	}
	defer rows.Close()

	// 取得した記事を１つ１つレスポンスとしてServer Streamingで返す
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
