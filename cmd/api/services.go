package main

import (
	"github.com/agung96tm/golearn-packages/internal/models"
)

func (app application) articleServiceCreate(req *ArticleCreateRequest) (*ArticleResponse, error) {
	var article models.Article
	if err := req.Validate(&article); err != nil {
		return nil, err
	}

	// save to db
	models.ArticleData = append(models.ArticleData, &article)

	// task
	app.runEmailArticleCreateTask(article.ID)

	return &ArticleResponse{
		ID:    article.ID,
		Title: article.Title,
		Body:  article.Body,
	}, nil
}
