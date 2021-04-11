package repository

import (
	"context"
	"database/sql"

	"github.com/k88t76/GraphQL-gRPC-demo/article/pb"
	_ "github.com/mattn/go-sqlite3"
)

type Repository interface {
	InsertArticle(ctx context.Context, input *pb.ArticleInput) (int64, error)
	SelectArticleByID(ctx context.Context, id int64) (*pb.Article, error)
	UpdateArticle(ctx context.Context, id int64, input *pb.ArticleInput) error
	DeleteArticle(ctx context.Context, id int64) error
	SelectAllArticles() (*sql.Rows, error)
}

type sqliteRepo struct {
	db *sql.DB
}

func NewsqliteRepo() (Repository, error) {
	db, err := sql.Open("sqlite3", "./article/article.sql")
	if err != nil {
		return nil, err
	}

	// articlesテーブルを作成
	cmd := `CREATE TABLE IF NOT EXISTS articles(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		author STRING,
		title STRING,
		content STRING)`

	_, err = db.Exec(cmd)
	if err != nil {
		return nil, err
	}
	return &sqliteRepo{db}, nil
}

func (r *sqliteRepo) InsertArticle(ctx context.Context, input *pb.ArticleInput) (int64, error) {
	// Inputの内容(Author, Title, Content)をarticlesテーブルにINSERT
	cmd := "INSERT INTO articles(author, title, content) VALUES (?, ?, ?)"
	result, err := r.db.Exec(cmd, input.Author, input.Title, input.Content)
	if err != nil {
		return 0, err
	}

	// INSERTした記事のIDを取得
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// INSERTした記事のIDを返す
	return id, nil
}

func (r *sqliteRepo) SelectArticleByID(ctx context.Context, id int64) (*pb.Article, error) {
	// 該当IDの記事をSELECT
	cmd := "SELECT * FROM articles WHERE id = ?"
	row := r.db.QueryRow(cmd, id)
	var a pb.Article

	// SELECTした記事の内容を読み取る
	err := row.Scan(&a.Id, &a.Author, &a.Title, &a.Content)
	if err != nil {
		return nil, err
	}

	// SELECTした記事を返す
	return &pb.Article{
		Id:      a.Id,
		Author:  a.Author,
		Title:   a.Title,
		Content: a.Content,
	}, nil
}

func (r *sqliteRepo) UpdateArticle(ctx context.Context, id int64, input *pb.ArticleInput) error {
	// 該当IDのAuthor, Title, ContentをUPDATE
	cmd := "UPDATE articles SET author = ?, title = ?, content = ? WHERE id = ?"
	_, err := r.db.Exec(cmd, input.Author, input.Title, input.Content, id)
	if err != nil {
		return err
	}

	// errorがなければ返り値なし
	return nil
}

func (r *sqliteRepo) DeleteArticle(ctx context.Context, id int64) error {
	// 該当IDの記事をDELETE
	cmd := "DELETE FROM articles WHERE id = ?"
	_, err := r.db.Exec(cmd, id)
	if err != nil {
		return err
	}

	// errorがなければ返り値なし
	return nil
}

func (r *sqliteRepo) SelectAllArticles() (*sql.Rows, error) {
	// articlesテーブルの記事を全取得
	cmd := "SELECT * FROM articles"
	rows, err := r.db.Query(cmd)
	if err != nil {
		return nil, err
	}

	// 全取得した記事を*sql.Rowsの形で返す
	return rows, nil
}
