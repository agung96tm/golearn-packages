package main

import (
	"context"
	"fmt"
	"github.com/agung96tm/golearn-packages/constants"
	"github.com/urfave/negroni"
	"net/http"
)

func (app application) coreMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		trxHandler := app.db.ORM.Begin()

		defer func() {
			if r := recover(); r != nil {
				trxHandler.Rollback()
			}
		}()

		ctx := context.WithValue(r.Context(), constants.DBTransaction, trxHandler)
		r = r.WithContext(ctx)

		lrw := negroni.NewResponseWriter(w)
		next.ServeHTTP(lrw, r)

		statusCode := lrw.Status()
		if statusCode >= 400 {
			trxHandler.Rollback()
		} else {
			if err := trxHandler.Commit().Error; err != nil {
				app.errorLog.Fatal(fmt.Sprintf("[database] transaction commit error: %v", err))
			}
		}
	})
}
