package models

import (
	"github.com/agung96tm/golearn-packages/lib"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title string
	Body  string
}

type ArticleModel struct {
	DB lib.Database
}

func (m ArticleModel) WithTrx(db *gorm.DB) ArticleModel {
	m.DB.ORM = db
	return m
}

func (m ArticleModel) Query() ([]*Article, error) {
	articles := make([]*Article, 0)
	err := m.DB.ORM.Model(&Article{}).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, err
}

func (m ArticleModel) Get(id uint) (*Article, error) {
	user := new(Article)
	err := m.DB.ORM.Model(user).Where("id=?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m ArticleModel) Create(article *Article) error {
	err := m.DB.ORM.Model(article).Create(article).Error
	if err != nil {
		return err
	}
	return nil
}

func (m ArticleModel) Update(article *Article) error {
	err := m.DB.ORM.Model(article).Updates(article).Error
	if err != nil {
		return err
	}
	return nil
}

func (m ArticleModel) Delete(article *Article) error {
	err := m.DB.ORM.Model(article).Delete(article).Error
	if err != nil {
		return err
	}
	return nil
}
