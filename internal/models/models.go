package models

import (
	"github.com/agung96tm/golearn-packages/lib"
)

type Models struct {
	Article ArticleModel
}

func NewModels(db lib.Database) Models {
	return Models{
		Article: ArticleModel{DB: db},
	}
}
