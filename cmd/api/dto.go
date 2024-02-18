package main

type ArticleCreateRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type ArticleUpdateRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type ArticleResponse struct {
	ID    uint   `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}
