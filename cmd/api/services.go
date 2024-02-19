package main

import (
	"errors"
	"github.com/agung96tm/golearn-packages/internal/models"
)

func (app application) articleServiceGetAll() []*ArticleResponse {
	resp := make([]*ArticleResponse, 0)

	articles, _ := app.models.Article.Query()
	for _, d := range articles {
		resp = append(resp, &ArticleResponse{
			ID:    d.ID,
			Title: d.Title,
			Body:  d.Body,
		})
	}

	return resp
}

func (app application) articleServiceGet(id uint) (*ArticleResponse, error) {
	if id <= 0 {
		return nil, errors.New("article not found")
	}

	article, err := app.models.Article.Get(id)
	if err != nil {
		return nil, err
	}

	resp := &ArticleResponse{
		ID:    article.ID,
		Title: article.Title,
		Body:  article.Body,
	}

	return resp, nil
}

/* ----------------------
Example with DB Transaction:

func (app application) articleServiceCreate(trxHandler *gorm.DB, req *ArticleCreateRequest) (*ArticleResponse, error) {
	...

	if err := app.models.Article.WithTrx(trxHandler).Create(&article); err != nil {
		return nil, err
	}

	...
}
---------------------- */

func (app application) articleServiceCreate(req *ArticleCreateRequest) (*ArticleResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	article := models.Article{
		Title: req.Title,
		Body:  req.Body,
	}

	if err := app.models.Article.Create(&article); err != nil {
		return nil, err
	}

	return &ArticleResponse{
		ID:    article.ID,
		Title: article.Title,
		Body:  article.Body,
	}, nil
}

func (app application) articleServiceUpdate(id uint, req *ArticleUpdateRequest) (*ArticleResponse, error) {
	article, err := app.models.Article.Get(id)
	if err != nil {
		return nil, err
	}

	if err := req.Validate(article); err != nil {
		return nil, err
	}

	err = app.models.Article.Update(article)
	if err != nil {
		return nil, err
	}

	return &ArticleResponse{
		ID:    article.ID,
		Title: article.Title,
		Body:  article.Body,
	}, nil
}

func (app application) articleServiceDelete(id uint) error {
	article, err := app.models.Article.Get(id)
	if err != nil {
		return err
	}

	err = app.models.Article.Delete(article)
	if err != nil {
		return err
	}

	return nil
}
