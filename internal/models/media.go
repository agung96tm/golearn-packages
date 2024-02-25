package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/agung96tm/golearn-packages/lib"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"
)

type MediaModel struct {
	DB      *lib.Database
	Storage *lib.Storage
}

func (m MediaModel) Get(id uint) (*Media, error) {
	query := `SELECT id, name, path FROM medias where id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var media Media
	err := m.DB.ORM.QueryRowContext(ctx, query, id).Scan(
		&media.ID,
		&media.Name,
		&media.Path,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &media, nil
}

func (m MediaModel) Create(media *Media) error {
	query := `
		INSERT INTO medias (name, path, upload_as) 
		VALUES ($1, $2, $3)
		RETURNING id
	`
	args := []interface{}{media.Name, media.Path, media.UploadAs}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.ORM.QueryRowContext(ctx, query, args...).Scan(&media.ID)
	if err != nil {
		return err
	}
	return nil
}

func (m MediaModel) Upload(fil *multipart.FileHeader) (*Media, error) {
	fileOpen, err := fil.Open()
	if err != nil {
		return nil, err
	}

	fileExt := filepath.Ext(fil.Filename)
	newFileName := strings.ReplaceAll(
		strings.ToLower(
			strings.TrimSuffix(filepath.Base(fil.Filename), fileExt),
		), " ", "-") + "-" + fmt.Sprintf(
		"%v", time.Now().Unix(),
	) + fileExt

	media := Media{
		Name:     newFileName,
		UploadAs: fil.Header.Get("Content-type"),
		Size:     fil.Size / 1024,
	}

	pathPut, err := m.Storage.Put(newFileName, fileOpen)
	if err != nil {
		return nil, err
	}
	media.Path = pathPut

	if err := m.Create(&media); err != nil {
		return nil, err
	}

	return &media, nil
}

func (m MediaModel) Delete(id uint) error {
	if id < 1 {
		return ErrNotFound
	}

	query := `DELETE FROM medias WHERE id = $1 RETURNING path`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var fullPath string
	err := m.DB.ORM.QueryRowContext(ctx, query, id).Scan(&fullPath)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrNotFound
		}
		return err
	}

	err = m.Storage.Delete(fullPath)
	if err != nil {
		return err
	}
	return nil
}

func (m MediaModel) DeleteAll(medias []*Media) {
	for _, media := range medias {
		err := m.Delete(media.ID)
		if err != nil {
		}
	}
}

type Media struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	UploadAs string `json:"upload_as"`
	Path     string `json:"path"`
}
