package main

import (
	"github.com/agung96tm/golearn-packages/internal/models"
	"mime/multipart"
)

func (app application) articleServiceGetAll() ([]*ArticleResponse, error) {
	articles, err := app.models.Article.GetAll()
	if err != nil {
		return nil, err
	}

	var resp []*ArticleResponse
	for _, article := range articles {
		resp = append(resp, &ArticleResponse{
			ID:    article.ID,
			Title: article.Title,
			Body:  article.Body,
		})
	}

	return resp, nil
}

func (app application) articleServiceGet(id uint) (*ArticleResponse, error) {
	article, err := app.models.Article.Get(id)
	if err != nil {
		return nil, err
	}

	return &ArticleResponse{
		ID:    article.ID,
		Title: article.Title,
		Body:  article.Body,
	}, nil
}

func (app application) articleServiceCreate(req *ArticleCreateRequest) (*ArticleResponse, error) {
	var article models.Article
	if err := req.Validate(&article); err != nil {
		return nil, err
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

	if err := app.models.Article.Update(article); err != nil {
		return nil, err
	}

	return &ArticleResponse{
		ID:    article.ID,
		Title: article.Title,
		Body:  article.Body,
	}, nil
}

func (app application) articleServiceDelete(id uint) error {
	err := app.models.Article.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (app application) articleServiceUpload(fils []*multipart.FileHeader) ([]*MediaResponse, error) {
	var medias []*models.Media
	for i := range fils {
		media, err := app.models.Media.Upload(fils[i])
		if err != nil {
			app.models.Media.DeleteAll(medias)
			return nil, err
		}
		medias = append(medias, media)
	}

	var resp []*MediaResponse
	for _, m := range medias {
		resp = append(resp, &MediaResponse{
			ID:   m.ID,
			Name: m.Name,
			Path: m.Path,
		})
	}
	return resp, nil
}
