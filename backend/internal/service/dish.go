package service

import (
	"fmt"

	"github.com/Zrossiz/Home_menu/backend/internal/dto"
	"github.com/Zrossiz/Home_menu/backend/internal/helpers"
)

type DishService struct {
	attachmentService   AttachmentService
	dishPostgresStorage DishPostgresStorage
}

type DishPostgresStorage interface {
	Create(payload dto.CreateDishDTO) error
	GetAllByCategory(categoryID int) ([]dto.DishDTO, error)
	Update(dishID int, payload dto.UpdateDishDTO) error
	Delete(dishID int) error
	GetOne(dishID int) (*dto.DishDTO, error)
	Find(search string) ([]dto.DishDTO, error)
}

func NewDishService(
	dishPostgresStorage DishPostgresStorage,
	attachmentService AttachmentService,
) *DishService {
	return &DishService{
		dishPostgresStorage: dishPostgresStorage,
		attachmentService:   attachmentService,
	}
}

func (d *DishService) Create(payload dto.CreateDishDTO) error {
	return d.dishPostgresStorage.Create(payload)
}

func (d *DishService) Update(dishID int, payload dto.UpdateDishDTO) error {
	return d.dishPostgresStorage.Update(dishID, payload)
}

func (d *DishService) Delete(dishID int) error {
	err := d.attachmentService.DeleteAllByDish(dishID)
	if err != nil {
		return fmt.Errorf("error remove images from filepath: %v", err)
	}

	err = d.dishPostgresStorage.Delete(dishID)
	if err != nil {
		return fmt.Errorf("error delete dish: %v", err)
	}

	return nil
}

func (d *DishService) GetAllByCategory(categoryID int) ([]dto.DishDTO, error) {
	return d.dishPostgresStorage.GetAllByCategory(categoryID)
}

func (d *DishService) GetOne(dishID int) (*dto.DishDTO, error) {
	return d.dishPostgresStorage.GetOne(dishID)
}

func (d *DishService) Find(search string) ([]dto.DishDTO, error) {
	return d.dishPostgresStorage.Find(search)
}

func (d *DishService) GetPublicPathsForImages(attachments []dto.AttachmentDTO) []string {
	var result []string
	for _, attachment := range attachments {
		publicPath := helpers.GetPublicPath(attachment.Key)
		result = append(result, publicPath)
	}

	return result
}
