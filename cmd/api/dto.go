package main

import (
	"github.com/agung96tm/golearn-packages/internal/models"
	"github.com/agung96tm/golearn-packages/internal/validator"
)

type ArticleCreateRequest struct {
	Title               string `json:"title"`
	Body                string `json:"body"`
	validator.Validator `json:"-"`
}

func (r *ArticleCreateRequest) Validate(article *models.Article) error {
	if !r.IsValid() {
		return validator.ErrValidator{
			Fields:    r.ErrFields,
			NonFields: r.ErrNonFields,
		}
	}
	return nil
}

type ArticleUpdateRequest struct {
	Title               string `json:"title"`
	Body                string `json:"body"`
	validator.Validator `json:"-"`
}

func (r *ArticleUpdateRequest) Validate(article *models.Article) error {
	if !r.IsValid() {
		return validator.ErrValidator{
			Fields:    r.ErrFields,
			NonFields: r.ErrNonFields,
		}
	}
	return nil
}

type ArticleResponse struct {
	ID    uint   `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}
