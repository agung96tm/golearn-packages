package models

import (
	"errors"
	"github.com/agung96tm/golearn-packages/lib"
)

var ErrNotFound = errors.New("data not found")

type Models struct {
	Article ArticleModel
}

func NewModels(db lib.Database) Models {
	return Models{
		Article: ArticleModel{DB: db},
	}
}
