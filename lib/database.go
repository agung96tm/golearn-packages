package lib

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)

type Database struct {
	ORM *sql.DB
}

func NewDB(dsn string) (*Database, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(15 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return &Database{
		ORM: db,
	}, nil
}
