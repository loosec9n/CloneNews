package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (a *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)

	if a.debug { //logging in debug mode
		mux.Use(middleware.Logger)

	}

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := a.render(w, r, "index", nil)
		if err != nil {
			log.Fatalln(err)
		}
	})
	mux.Get("/comment", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("comments in the website"))
	})

	return mux
}
