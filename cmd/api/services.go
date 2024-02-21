package main

import (
	"errors"
	"github.com/agung96tm/golearn-packages/internal/models"
)

func (app application) articleServiceGetAll() []*models.Article {
	return models.ArticleData
}

func (app application) articleServiceGet(id uint) (*models.Article, error) {
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

func (app application) articleServiceCreate(req *ArticleCreateRequest) (*ArticleResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	var article models.Article
	id := uint(len(models.ArticleData) + 1)

	article.ID = id
	if req.Title != nil {
		article.Title = *req.Title
	}
	if req.Body != nil {
		article.Body = *req.Body
	}

	// save to db
	models.ArticleData = append(models.ArticleData, &article)

	return &ArticleResponse{
		ID:    article.ID,
		Title: article.Title,
		Body:  article.Body,
	}, nil
}

func (app application) articleServiceUpdate(id uint, data ArticleUpdateRequest) (*ArticleResponse, error) {
	if id <= 0 {
		return nil, models.ErrNotFound
	}

	article, err := app.articleServiceGet(id)
	if err != nil {
		return nil, err
	}

	if data.Title != nil {
		article.Title = *data.Title
	}
	if data.Body != nil {
		article.Body = *data.Body
	}

	// update db
	for i, articleDb := range models.ArticleData {
		if articleDb.ID == id {
			models.ArticleData[i] = article
		}
	}

	return &ArticleResponse{
		ID:    article.ID,
		Title: article.Title,
		Body:  article.Body,
	}, nil
}
