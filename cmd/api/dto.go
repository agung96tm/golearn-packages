package main

import "github.com/agung96tm/golearn-packages/internal/validator"

type ArticleCreateRequest struct {
	Title               *string `json:"title"`
	Body                *string `json:"body"`
	validator.Validator `json:"-"`
}

func (v *ArticleCreateRequest) Validate() error {
	if v.Title == nil || *v.Title == "" {
		v.AddErrField("title", "the field is required")
	}
	if v.Body == nil || *v.Body == "" {
		v.AddErrField("body", "the field is required")
	}

	if !v.IsValid() {
		return validator.ErrValidator{
			Fields:    v.ErrFields,
			NonFields: v.ErrNonFields,
		}
	}
	return nil
}

type ArticleUpdateRequest struct {
	Title *string `json:"title"`
	Body  *string `json:"body"`
}

type ArticleResponse struct {
	ID    uint   `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}
