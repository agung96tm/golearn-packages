package main

import (
	"github.com/agung96tm/golearn-packages/internal/models"
)

func (app application) articleServiceGetAll() []*models.Article {
	return models.ArticleData
}

func (app application) articleServiceCreate(articleForm *ArticleForm) (*models.Article, error) {
	var article models.Article
	if err := articleForm.Validate(&article); err != nil {
		return nil, err
	}

	// save to db
	models.ArticleData = append(models.ArticleData, &article)

	// task
	app.runEmailArticleCreateTask(article.ID)

	return &article, nil
}
