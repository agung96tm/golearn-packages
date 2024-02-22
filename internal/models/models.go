package models

import (
	"errors"
	"github.com/agung96tm/golearn-packages/lib"
)

var (
	ErrNotFound = errors.New("data not found")
)

type Model struct {
	Media   MediaModel
	Article ArticleModel
}

func New(db *lib.Database) Model {
	return Model{
		Media:   MediaModel{DB: db},
		Article: ArticleModel{DB: db},
	}
}
