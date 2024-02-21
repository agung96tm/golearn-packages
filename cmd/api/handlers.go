package main

import (
	"net/http"
)

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

	err = app.writeJSON(w, http.StatusCreated, resp, nil)
	if err != nil {
		app.serverErrorResponse(w, err)
	}
}
