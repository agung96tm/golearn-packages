package main

import (
	"github.com/agung96tm/golearn-packages/internal/models"
)

func (app *application) ArticleServiceGetAll() ([]*models.Article, error) {
	articles, err := app.models.Article.Query()
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (app *application) ArticleServiceGet(id uint) (*models.Article, error) {
	article, err := app.models.Article.Get(id)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (app *application) ArticleServiceCreate(form *ArticleForm) (*models.Article, error) {
	if err := form.Validate(); err != nil {
		return nil, err
	}

	article := models.Article{
		Title: form.Title,
		Body:  form.Body,
	}

	err := app.models.Article.Create(&article)
	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (app *application) ArticleServiceUpdate(id uint, form *ArticleEditForm) (*models.Article, error) {
	article, err := app.models.Article.Get(id)
	if err != nil {
		return nil, models.ErrNotFound
	}

	if err := form.Validate(article); err != nil {
		return article, err
	}

	err = app.models.Article.Update(article)
	if err != nil {
		return article, err
	}

	return article, nil
}

func (app *application) ArticleServiceDelete(id uint) error {
	article, err := app.models.Article.Get(id)
	if err != nil {
		return models.ErrNotFound
	}

	err = app.models.Article.Delete(article)
	if err != nil {
		return err
	}

	return nil
}
