package handler

import "go.uber.org/zap"

type Handler struct {
	Category   CategoryHandler
	Dish       DishHandler
	Attachment AttachmentHandler
}

type Service struct {
	Category   CategoryService
	Dish       DishService
	Attachment AttachmentService
}

func New(service Service, logger *zap.Logger) *Handler {
	var handler Handler

	handler.Category = *NewCategoryHandler(service.Category, logger)
	handler.Dish = *NewDishHandler(service.Dish, service.Attachment, logger)
	handler.Attachment = *NewAttachmentHandler(service.Attachment, logger)

	return &handler
}
