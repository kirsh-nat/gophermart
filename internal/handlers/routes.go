package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func Routes(handler *URLHandler) *chi.Mux {

	r := chi.NewRouter()
	fmt.Print("hello")
	//r.Post("/api/user/register", http.HandlerFunc(Middleware(http.HandlerFunc(handler.Registration))))
	r.Post("/api/user/register", http.HandlerFunc(handler.Registration))
	// r.Get("/{id}", handler.Get)
	// r.Post("/api/shorten", http.HandlerFunc(Middleware(http.HandlerFunc(handler.GetAPIShorten))))
	// r.Get("/ping", http.HandlerFunc(Middleware(http.HandlerFunc(handler.PingHandler))))
	// r.Post("/api/shorten/batch", http.HandlerFunc(Middleware(http.HandlerFunc(handler.AddBatch))))
	// r.Get("/api/user/urls", http.HandlerFunc(Middleware(http.HandlerFunc(handler.GetUserURLs))))
	// r.Delete("/api/user/urls", http.HandlerFunc(Middleware(http.HandlerFunc(handler.DeleteUserURLs))))

	return r
}
