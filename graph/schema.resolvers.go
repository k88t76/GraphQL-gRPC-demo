package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/k88t76/GraphQL-gRPC-demo/article/pb"
	"github.com/k88t76/GraphQL-gRPC-demo/graph/generated"
	"github.com/k88t76/GraphQL-gRPC-demo/graph/model"
)

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (m *mutationResolver) CreateArticle(ctx context.Context, input model.CreateInput) (*model.Article, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	article, err := m.Server.ArticleClient.CreateArticle(
		ctx,
		&pb.CreateInput{
			Author:  input.Author,
			Title:   input.Title,
			Content: input.Content,
		})
	if err != nil {
		return nil, err
	}
	return &model.Article{
		ID:      int(article.ID),
		Author:  article.Author,
		Title:   article.Title,
		Content: article.Content,
	}, nil
}

func (m *mutationResolver) UpdateArticle(ctx context.Context, input model.UpdateInput) (*model.Article, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	article, err := m.Server.ArticleClient.UpdateArticle(
		ctx,
		&pb.Article{
			Id:      int64(input.ID),
			Author:  input.Author,
			Title:   input.Title,
			Content: input.Content,
		})
	if err != nil {
		return nil, err
	}

	return &model.Article{
		ID:      int(article.ID),
		Author:  article.Author,
		Title:   article.Title,
		Content: article.Content,
	}, nil
}

func (m *mutationResolver) DeleteArticle(ctx context.Context, id int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	articleId, err := m.Server.ArticleClient.DeleteArticle(ctx, int64(id))
	if err != nil {
		return 0, err
	}
	return int(articleId), nil
}

type queryResolver struct{ *Resolver }

func (q *queryResolver) Article(ctx context.Context, id int) (*model.Article, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	article, err := q.Server.ArticleClient.ReadArticle(ctx, int64(id))
	if err != nil {
		return nil, err
	}
	return &model.Article{
		ID:      int(article.ID),
		Author:  article.Author,
		Title:   article.Title,
		Content: article.Content,
	}, nil
}

func (q *queryResolver) Articles(ctx context.Context) ([]*model.Article, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	articles, err := q.Server.ArticleClient.ListArticle(ctx)
	if err != nil {
		return nil, err
	}
	return articles, nil
}
