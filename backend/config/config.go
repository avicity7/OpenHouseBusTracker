package config

import (
	"context"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olahol/melody"
)

// To evaluate performance using single connection vs Pool in the future.
var Dbpool *pgxpool.Pool
var err error
var Melody *melody.Melody

func Connect(r *chi.Mux) {
	Melody = melody.New()
	r.Get("/ws", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { Melody.HandleRequest(w, r) }))
	Dbpool, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		println(err)
		println("unable to connect")
	}
}
