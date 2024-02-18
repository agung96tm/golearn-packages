package main

import (
	"github.com/agung96tm/golearn-packages/internal/models"
)

func (app application) ArticleServiceGetAll() []*models.Article {
	return models.ArticleData
}

func (app application) ArticleServiceGet(id uint) (*models.Article, error) {
	return nil, nil
}

func (app application) ArticleServiceCreate(data ArticleCreateRequest) (*ArticleResponse, error) {
	return nil, nil
}

func (app application) ArticleServiceUpdate(id uint, data ArticleUpdateRequest) (*ArticleResponse, error) {
	return nil, nil
}
