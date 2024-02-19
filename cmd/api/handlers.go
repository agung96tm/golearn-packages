package main

import (
	"net/http"
)

func (app application) articleList(w http.ResponseWriter, r *http.Request) {
	resp := app.articleServiceGetAll()

	err := app.writeJSON(w, http.StatusOK, resp, nil)
	if err != nil {
		app.serverErrorResponse(w, err)
	}
}

func (app application) articleDetail(w http.ResponseWriter, r *http.Request) {
	id, _ := app.readIDParam(r)
	resp, err := app.articleServiceGet(id)
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

	/*
		-- Example with DB Transaction --
		trxHandle := r.Context().Value(constants.DBTransaction).(*gorm.DB)
		resp, err := app.articleServiceCreate(trxHandle, req)
	*/

	resp, err := app.articleServiceCreate(req)
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

	resp, err := app.articleServiceUpdate(id, &req)
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

	err = app.articleServiceDelete(id)
	if err != nil {
		app.badRequestResponse(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusNoContent, nil, nil)
	if err != nil {
		app.serverErrorResponse(w, err)
	}
}
