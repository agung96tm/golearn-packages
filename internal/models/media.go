package models

import (
	"github.com/agung96tm/golearn-packages/lib"
)

type MediaModel struct {
	DB *lib.Database
}

type Media struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	UploadTo string `json:"upload_to"`
}
