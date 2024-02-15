package main

import (
	"github.com/agung96tm/golearn-packages/internal/models"
	"net/http"
)

func (app application) articleList(w http.ResponseWriter, r *http.Request) {
	err := app.writeJSON(w, http.StatusOK, models.ArticleData, nil)
	if err != nil {
		app.serverErrorResponse(w, err)
		return
	}
}

func (app application) articleCreate(w http.ResponseWriter, r *http.Request) {
	var req ArticleCreateRequest

	err := app.readJSON(w, r, &req)
	if err != nil {
		app.badRequestResponse(w, err)
		return
	}

	/* SAVE ARTICLES */
	models.ArticleData = append(models.ArticleData, models.Article{
		Title: req.Title,
		Body:  req.Body,
	})

	err = app.writeJSON(w, http.StatusCreated, req, nil)
	if err != nil {
		app.serverErrorResponse(w, err)
	}
}
