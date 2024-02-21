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
	for _, d := range models.ArticleData {
		if d.ID == id {
			article = d
		}
	}
	if article == nil {
		return nil, errors.New("article not found")
	}
	return article, nil
}

func (app application) articleServiceCreate() {
	return
}

func (app application) articleServiceUpdate() {
	return
}
