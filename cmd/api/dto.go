package main

import "github.com/agung96tm/golearn-packages/internal/validator"

type ArticleCreateRequest struct {
	Title               string `json:"title"`
	Body                string `json:"body"`
	validator.Validator `json:"-"`
}

type ArticleUpdateRequest struct {
	Title               string `json:"title"`
	Body                string `json:"body"`
	validator.Validator `json:"-"`
}

type ArticleResponse struct {
	ID    uint   `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}
