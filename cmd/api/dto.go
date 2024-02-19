package main

import (
	"github.com/agung96tm/golearn-packages/internal/models"
	"github.com/agung96tm/golearn-packages/internal/validator"
)

type ArticleCreateRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`

	validator.Validator `json:"-"`
}

func (a ArticleCreateRequest) Validate() error {
	if a.Title == "" {
		a.AddErrField("title", "the field is required")
	}
	if a.Body == "" {
		a.AddErrField("body", "the field is required")
	}

	if !a.IsValid() {
		return validator.ErrValidator{
			Fields:    a.ErrFields,
			NonFields: a.ErrNonFields,
		}
	}
	return nil
}

type ArticleUpdateRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`

	validator.Validator `json:"-"`
}

func (a ArticleUpdateRequest) Validate(article *models.Article) error {
	//if a.Title == "" {
	//	a.AddErrField("title", "the field is required")
	//}
	//if a.Body == "" {
	//	a.AddErrField("body", "the field is required")
	//}
	if !a.IsValid() {
		return validator.ErrValidator{
			Fields:    a.ErrFields,
			NonFields: a.ErrNonFields,
		}
	}

	if a.Title != article.Title {
		article.Title = a.Title
	}
	if a.Body != article.Body {
		article.Body = a.Body
	}

	return nil
}

type ArticleResponse struct {
	ID    uint   `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}
