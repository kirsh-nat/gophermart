package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Routes(handler *URLHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/api/user/register", http.HandlerFunc(handler.Registration))
	r.Post("/api/user/login", http.HandlerFunc(handler.Authentication))
	r.Post("/api/user/orders", http.HandlerFunc(handler.CreateOrder))
	r.Get("/api/user/orders", http.HandlerFunc(handler.GetOrders))
	r.Get("/api/user/balance", http.HandlerFunc(handler.GetBalance))
	r.Get("/api/user/withdrawals", http.HandlerFunc(handler.GetDrafts))
	r.Post("/api/user/balance/withdraw", http.HandlerFunc(handler.CreateDraft))

	// r.Get("/{id}", handler.Get)
	// r.Post("/api/shorten", http.HandlerFunc(Middleware(http.HandlerFunc(handler.GetAPIShorten))))
	// r.Get("/ping", http.HandlerFunc(Middleware(http.HandlerFunc(handler.PingHandler))))
	// r.Post("/api/shorten/batch", http.HandlerFunc(Middleware(http.HandlerFunc(handler.AddBatch))))
	// r.Get("/api/user/urls", http.HandlerFunc(Middleware(http.HandlerFunc(handler.GetUserURLs))))
	// r.Delete("/api/user/urls", http.HandlerFunc(Middleware(http.HandlerFunc(handler.DeleteUserURLs))))

	return r
}
