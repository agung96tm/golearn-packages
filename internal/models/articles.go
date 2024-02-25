package models

import (
	"context"
	"database/sql"
	"errors"
	"github.com/agung96tm/golearn-packages/lib"
	"time"
)

type ArticleModel struct {
	DB *lib.Database
}

func (m ArticleModel) GetAll() (Articles, error) {
	query := `
		SELECT articles.id, 
			   articles.title, 
			   articles.body, 
			   medias.id, 
			   medias.name, 
			   medias.path
		FROM articles
		LEFT JOIN medias ON articles.image_id = medias.id
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.ORM.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var articles Articles
	for rows.Next() {
		var article Article

		err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Body,
			&article.Image.ID,
			&article.Image.Name,
			&article.Image.Path,
		)
		if err != nil {
			return nil, err
		}

		articles = append(articles, &article)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}

func (m ArticleModel) Get(id uint) (*Article, error) {
	if id <= 0 {
		return nil, ErrNotFound
	}

	query := `
		SELECT articles.id, 
			   articles.title, 
			   articles.body, 
			   medias.id, 
			   medias.name, 
			   medias.path
		FROM articles
		LEFT JOIN medias ON articles.image_id = medias.id
		WHERE articles.id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var article Article
	err := m.DB.ORM.QueryRowContext(ctx, query, id).Scan(
		&article.ID,
		&article.Title,
		&article.Body,
		&article.Image.ID,
		&article.Image.Name,
		&article.Image.Path,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &article, nil
}

func (m ArticleModel) Create(article *Article) error {
	query := `
		INSERT INTO articles (title, body, image_id) 
		VALUES ($1, $2, $3)
		RETURNING id
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{article.Title, article.Body, article.ImageID}

	return m.DB.ORM.QueryRowContext(ctx, query, args...).Scan(
		&article.ID,
	)
}

func (m ArticleModel) Update(article *Article) error {
	query := `
		UPDATE articles
		SET title = $1, body = $2
		WHERE id = $3
		RETURNING id
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{article.Title, article.Body, article.ID}
	err := m.DB.ORM.QueryRowContext(ctx, query, args...).Scan(&article.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrNotFound
		default:
			return err
		}
	}

	return nil
}

func (m ArticleModel) Delete(id uint) error {
	if id < 1 {
		return ErrNotFound
	}

	query := `DELETE FROM articles WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ORM.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

type MediaInArticle struct {
	ID   sql.NullInt64
	Name sql.NullString
	Path sql.NullString
}

type Article struct {
	ID      uint           `json:"id"`
	Title   string         `json:"title"`
	Body    string         `json:"body"`
	ImageID uint           `json:"image_id"`
	Image   MediaInArticle `json:"image"`
}

type Articles []*Article
