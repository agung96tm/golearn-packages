package main

import (
	"net/http"
)

func (app application) articleList(w http.ResponseWriter, r *http.Request) {
	articles := app.ArticleServiceGetAll()

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

	article, err := app.ArticleServiceGet(id)
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
	err := app.readJSON(w, r, &req)
	if err != nil {
		app.badRequestResponse(w, err)
		return
	}

	resp, err := app.ArticleServiceCreate(req)
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
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var req ArticleUpdateRequest
	err = app.readJSON(w, r, &req)
	if err != nil {
		app.badRequestResponse(w, err)
		return
	}

	resp, err := app.ArticleServiceUpdate(id, req)
	if err != nil {
		app.badRequestResponse(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, resp, nil)
	if err != nil {
		app.serverErrorResponse(w, err)
	}
}
