package main

import (
	"github.com/agung96tm/golearn-packages/internal/models"
)

func (app application) articleServiceGetAll() (models.Articles, error) {
	articles, err := app.models.Article.GetAll()
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (app application) articleServiceGet(id uint) (*models.Article, error) {
	article, err := app.models.Article.Get(id)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (app application) articleServiceCreate(articleForm *ArticleForm) (*models.Article, error) {
	var article models.Article

	if err := articleForm.Validate(&article); err != nil {
		return nil, err
	}
	if err := app.models.Article.Create(&article); err != nil {
		return nil, err
	}

	return &article, nil
}

func (app application) articleServiceUpdate(id uint, articleForm *ArticleEditForm) (*models.Article, error) {
	article, err := app.models.Article.Get(id)
	if err != nil {
		return nil, err
	}

	if err := articleForm.Validate(article); err != nil {
		return nil, err
	}
	if err := app.models.Article.Update(article); err != nil {
		return nil, err
	}

	return article, err
}

func (app application) articleServiceDelete(id uint) error {
	err := app.models.Article.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
