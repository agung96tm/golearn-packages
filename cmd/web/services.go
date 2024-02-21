package main

import (
	"errors"
	"github.com/agung96tm/golearn-packages/internal/models"
)

func (app application) articleServiceGetAll() []*models.Article {
	return models.ArticleData
}

func (app application) articleServiceGet(id uint) (*models.Article, error) {
	if id <= 0 {
		return nil, models.ErrNotFound
	}

	var article *models.Article
	for _, articleDb := range models.ArticleData {
		if articleDb.ID == id {
			article = articleDb
		}
	}
	if article == nil {
		return nil, errors.New("article not found")
	}
	return article, nil
}

func (app application) articleServiceCreate(form *ArticleForm) error {
	var article models.Article
	if err := form.Validate(&article); err != nil {
		return err
	}

	// save to DB
	models.ArticleData = append(models.ArticleData, &models.Article{
		ID:    uint(len(models.ArticleData) + 1),
		Title: *form.Title,
		Body:  *form.Body,
	})

	return nil
}

func (app application) articleServiceUpdate(article *models.Article, articleForm *ArticleEditForm) error {
	if err := articleForm.Validate(article); err != nil {
		return err
	}

	// save to DB
	for i, articleDb := range models.ArticleData {
		if articleDb.ID == article.ID {
			models.ArticleData[i] = article
		}
	}

	return nil
}
