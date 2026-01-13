// Package config sets up the DB connection, WebSocket server and in-memory cache.
package config

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olahol/melody"
	"github.com/patrickmn/go-cache"
)

var Dbpool *pgxpool.Pool
var err error
var Melody *melody.Melody
var Cache *cache.Cache

func Connect(r *chi.Mux) {
	Melody = melody.New()
	r.Get("/ws", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { Melody.HandleRequest(w, r) }))
	Dbpool, err = pgxpool.New(context.Background(), fmt.Sprintf("postgresql://%s:%s@%s:%s/fyp", os.Getenv("DB_USER"), os.Getenv("DB_SECRET"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT")))
	if err != nil {
		println(err)
		println("unable to connect")
	}
	Cache = cache.New(5*time.Minute, 10*time.Minute)
}
