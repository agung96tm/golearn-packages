package main

import (
	"errors"
	"github.com/agung96tm/golearn-packages/internal/models"
	"net/http"
)

func (app application) articleList(w http.ResponseWriter, r *http.Request) {
	articles := app.articleServiceGetAll()

	err := app.writeJSON(w, http.StatusOK, articles, nil)
	if err != nil {
		app.serverErrorResponse(w, err)
		return
	}
}

func (app application) articleDetail(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, "")
		return
	}

	article, err := app.articleServiceGet(id)
	if err != nil {
		app.notFoundResponse(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, article, nil)
	if err != nil {
		app.serverErrorResponse(w, err)
	}
}

func (app application) articleCreate(w http.ResponseWriter, r *http.Request) {
	var req ArticleCreateRequest
	if err := app.readJSON(w, r, &req); err != nil {
		app.badRequestResponse(w, err)
		return
	}

	resp, err := app.articleServiceCreate(&req)
	if err != nil {
		app.badRequestResponse(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, resp, nil)
	if err != nil {
		app.serverErrorResponse(w, err)
	}
}

func (app application) articleUpdate(w http.ResponseWriter, r *http.Request) {
	var req ArticleUpdateRequest
	if err := app.readJSON(w, r, &req); err != nil {
		app.badRequestResponse(w, err)
		return
	}

	id, _ := app.readIDParam(r)
	resp, err := app.articleServiceUpdate(id, req)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			app.notFoundResponse(w, err)
		default:
			app.badRequestResponse(w, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, resp, nil)
	if err != nil {
		app.serverErrorResponse(w, err)
	}
}
