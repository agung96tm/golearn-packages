package main

import (
	"fmt"
	"net/http"
)

func (app application) articleList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list"))
}

func (app application) articleDetail(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("detail"))
}

func (app application) articleCreate(w http.ResponseWriter, r *http.Request) {
	err := app.runEmailDeliveryTask(uint(1))
	if err != nil {
		fmt.Println("error", err)
	}

	w.Write([]byte("create"))
}

func (app application) articleUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update"))
}
