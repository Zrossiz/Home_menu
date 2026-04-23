package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type CategoryRouter struct {
	handler CategoryHandler
}

type CategoryHandler interface {
	Create(rw http.ResponseWriter, r *http.Request)
	Delete(rw http.ResponseWriter, r *http.Request)
	GetAll(rw http.ResponseWriter, r *http.Request)
	Update(rw http.ResponseWriter, r *http.Request)
}

func NewCategoryRouter(categoryHandler CategoryHandler) *CategoryRouter {
	return &CategoryRouter{
		handler: categoryHandler,
	}
}

func RegisterCategoryRoutes(r chi.Router, handler CategoryHandler) {
	r.Route("/api/category", func(r chi.Router) {
		r.Post("/", handler.Create)
		r.Get("/", handler.GetAll)
		r.Post("/{categoryID}", handler.Update)
		r.Delete("/{categoryID}", handler.Delete)
	})
}
