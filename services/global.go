package service

import (
	"context"
	"database/sql"
	"iqbvl/go-graphql-intro/models"
	"log"
	"net/http"
)

var dbConn *sql.DB

func OpenDB() (*sql.DB, error) {
	conn, err := sql.Open("sqlserver", models.BuildConnString())
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	dbConn = conn

	dbConn.SetMaxOpenConns(5)
	dbConn.SetMaxIdleConns(1)
	dbConn.SetConnMaxLifetime(0)
	return dbConn, err
}

// CONTEXT
func AddContext(ctx context.Context, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

type ContextKey string

func (c ContextKey) String() string {
	return string(c)
}

var (
	// Ctx UserID
	CtxUserID  = ContextKey("user_id")
	CtxIsAuth  = ContextKey("is_authorized")
	CtxIsAdmin = ContextKey("is_admin")
)

func ArrayContains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
