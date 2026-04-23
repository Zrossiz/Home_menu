package service

import (
	"github.com/Zrossiz/Home_menu/backend/internal/dto"
)

type CategoryService struct {
	categoryPostgresStorage CategoryPostgresStorage
}

type CategoryPostgresStorage interface {
	Create(dto.CreateCategoryDTO) error
	Update(categoryID int, payload dto.UpdateCategoryDTO) error
	Delete(categoryID int) error
	GetAll() ([]dto.CategoryDTO, error)
}

func NewCategoryService(categoryPostgresStorage CategoryPostgresStorage) *CategoryService {
	return &CategoryService{
		categoryPostgresStorage: categoryPostgresStorage,
	}
}

func (c *CategoryService) Create(payload dto.CreateCategoryDTO) error {
	return c.categoryPostgresStorage.Create(payload)
}

func (c *CategoryService) Update(categoryID int, payload dto.UpdateCategoryDTO) error {
	return c.categoryPostgresStorage.Update(categoryID, payload)
}

func (c *CategoryService) Delete(categoryID int) error {
	return c.categoryPostgresStorage.Delete(categoryID)
}

func (c *CategoryService) GetAll() ([]dto.CategoryDTO, error) {
	return c.categoryPostgresStorage.GetAll()
}
