package dto

import "time"

type CreateCategoryDTO struct {
	Name string `json:"name"`
}

type UpdateCategoryDTO struct {
	Name string `json:"name"`
}

type CategoryDTO struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
