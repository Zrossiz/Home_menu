package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type DishRouter struct {
	handler DishHandler
}

type DishHandler interface {
	Create(rw http.ResponseWriter, r *http.Request)
	Delete(rw http.ResponseWriter, r *http.Request)
	GetAllByCategory(rw http.ResponseWriter, r *http.Request)
	GetOne(rw http.ResponseWriter, r *http.Request)
	Update(rw http.ResponseWriter, r *http.Request)
	Find(rw http.ResponseWriter, r *http.Request)
}

func NewDishRouter(dishHandler DishHandler) *DishRouter {
	return &DishRouter{
		handler: dishHandler,
	}
}

func RegisterDishRoutes(r chi.Router, handler DishHandler) {
	r.Route("/api/dish", func(r chi.Router) {
		r.Get("/search", handler.Find)
		r.Post("/", handler.Create)
		r.Post("/{dishID}", handler.Update)
		r.Get("/{dishID}", handler.GetOne)
		r.Delete("/{dishID}", handler.Delete)
		r.Get("/category/{categoryID}", handler.GetAllByCategory)
	})
}
