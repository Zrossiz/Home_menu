package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Category   CategoryHandler
	Dish       DishHandler
	Attachment AttachmentHandler
}

func New(handler Handler) http.Handler {
	r := chi.NewRouter()

	RegisterCategoryRoutes(r, handler.Category)
	RegisterDishRoutes(r, handler.Dish)
	RegisterAttachmentRoutes(r, handler.Attachment)

	return r
}
