package main

import (
	"github.com/agung96tm/golearn-packages/internal/models"
)

func (app application) articleServiceGetAll() []*models.Article {
	return models.ArticleData
}

func (app application) articleServiceGet(id uint) (*models.Article, error) {
	return models.ArticleData[0], nil
}

func (app application) articleServiceCreate(req *ArticleForm) (*models.Article, error) {
	return nil, nil
}

func (app application) articleServiceUpdate(id uint, req *ArticleEditForm) (*models.Article, error) {
	return nil, nil
}

func (app application) articleServiceDelete(id uint) error {
	return nil
}
