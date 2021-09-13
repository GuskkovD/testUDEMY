package main

import (
	"github.com/GuskkovD/go-course/pkg/config"
	"github.com/GuskkovD/go-course/pkg/handlers"
	"github.com/bmizerany/pat"
	"net/http"
)

//routes sets multiplexers
func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux
}
