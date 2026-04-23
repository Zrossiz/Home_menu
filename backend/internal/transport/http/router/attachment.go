package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AttachmentRouter struct {
	handler AttachmentHandler
}

type AttachmentHandler interface {
	Create(rw http.ResponseWriter, r *http.Request)
	GetOne(rw http.ResponseWriter, r *http.Request)
}

func NewAttachmentRouter(handler AttachmentHandler) *AttachmentRouter {
	return &AttachmentRouter{handler: handler}
}

func RegisterAttachmentRoutes(r chi.Router, handler AttachmentHandler) {
	r.Route("/api/attachment", func(r chi.Router) {
		r.Post("/dish/{dishID}", handler.Create)
		r.Get("/{key}", handler.GetOne)
	})
}
