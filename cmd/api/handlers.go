package main

import (
	"net/http"
)

func (app application) articleList(w http.ResponseWriter, r *http.Request) {
	resp := app.ArticleServiceGetAll()

	err := app.writeJSON(w, http.StatusOK, resp, nil)
	if err != nil {
		app.serverErrorResponse(w, err)
	}
}

func (app application) articleDetail(w http.ResponseWriter, r *http.Request) {
	id, _ := app.readIDParam(r)
	resp, err := app.ArticleServiceGet(id)
	if err != nil {
		app.notFoundResponse(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, resp, nil)
	if err != nil {
		app.serverErrorResponse(w, err)
	}
}

func (app application) articleCreate(w http.ResponseWriter, r *http.Request) {
	req := new(ArticleCreateRequest)

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

	err = app.writeJSON(w, http.StatusCreated, resp, nil)
	if err != nil {
		app.serverErrorResponse(w, err)
	}
}

func (app application) articleUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, err)
		return
	}

	var req ArticleUpdateRequest
	if err := app.readJSON(w, r, &req); err != nil {
		app.badRequestResponse(w, err)
		return
	}

	resp, err := app.ArticleServiceUpdate(id, &req)
	if err != nil {
		app.badRequestResponse(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, resp, nil)
	if err != nil {
		app.serverErrorResponse(w, err)
	}
}

func (app application) articleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, err)
		return
	}

	err = app.ArticleServiceDelete(id)
	if err != nil {
		app.badRequestResponse(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusNoContent, nil, nil)
	if err != nil {
		app.serverErrorResponse(w, err)
	}
}
