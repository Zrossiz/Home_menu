package service

import "go.uber.org/zap"

type Service struct {
	Category   CategoryService
	Dish       DishService
	Attachment AttachmentService
}

type Storage struct {
	CategoryPostgres   CategoryPostgresStorage
	DishPostgres       DishPostgresStorage
	AttachmentPostgres AttachmentPostgresStorage
}

func New(store Storage, log *zap.Logger) *Service {
	var svc Service

	svc.Category = *NewCategoryService(store.CategoryPostgres)
	svc.Attachment = *NewAttachmentService(store.AttachmentPostgres)
	svc.Dish = *NewDishService(store.DishPostgres, svc.Attachment)

	return &svc
}
