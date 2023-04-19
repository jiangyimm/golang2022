package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/test", Test)
	http.ListenAndServe(":3000", r)
}

func Test(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("123"))
}
