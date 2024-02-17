package main

import (
	"errors"
	"github.com/agung96tm/golearn-packages/internal/models"
)

func (app *application) ArticleServiceGetAll() []models.Article {
	return models.ArticleData
}

func (app *application) ArticleServiceGet(id uint) (*models.Article, error) {
	if id <= 0 {
		return nil, errors.New("article not found")
	}

	var article *models.Article
	for _, articleDb := range models.ArticleData {
		if articleDb.ID == id {
			article = &articleDb
		}
	}
	if article == nil {
		return nil, errors.New("article not found")
	}
	return article, nil
}

func (app *application) ArticleServiceCreate(form *ArticleForm) {
	var article ArticleForm

	if form.Title != nil {
		//form.SetErrField("Title", "Invalid Format Maybe")
		article.Title = form.Title
	}
	if form.Body != nil {
		article.Body = form.Body
	}

	//form.SetErrNonField("Invalid something")

	// save to DB
	models.ArticleData = append(models.ArticleData, models.Article{
		ID:    uint(len(models.ArticleData) + 1),
		Title: *form.Title,
		Body:  *form.Body,
	})
}

func (app *application) ArticleServiceUpdate(article *models.Article, form *ArticleEditForm) {
	if form.Title != nil {
		//form.SetErrField("Title", "Invalid Format Maybe")
		article.Title = *form.Title
	}
	if form.Body != nil {
		article.Body = *form.Body
	}

	//form.SetErrNonField("Invalid something")

	// save to DB
	for i, articleDb := range models.ArticleData {
		if articleDb.ID == article.ID {
			models.ArticleData[i] = *article
		}
	}
}
