package main

import (
	"github.com/agung96tm/golearn-packages/internal/models"
)

func (app application) articleServiceGetAll() []*models.Article {
	return models.ArticleData
}

func (app application) articleServiceGet(id uint) (*models.Article, error) {
	return nil, nil
}

func (app application) articleServiceCreate(data ArticleCreateRequest) (*ArticleResponse, error) {
	return nil, nil
}

func (app application) articleServiceUpdate(id uint, data ArticleUpdateRequest) (*ArticleResponse, error) {
	return nil, nil
}

func (app application) articleServiceDelete(id uint) error {
	return nil
}
