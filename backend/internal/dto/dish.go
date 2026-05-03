package dto

import "time"

type CreateDishDTO struct {
	Name        string `json:"name"`
	TimeToCook  int    `json:"time_to_cook"`
	Recipe      string `json:"recipe"`
	Description string `json:"description"`
	CategoryID  int    `json:"category_id"`
}

type UpdateDishDTO struct {
	Name        string `json:"name"`
	TimeToCook  int    `json:"time_to_cook"`
	Recipe      string `json:"recipe"`
	Description string `json:"description"`
}

type DishDTO struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	TimeToCook  int       `json:"time_to_cook"`
	Recipe      string    `json:"recipe"`
	Description string    `json:"description"`
	CategoryID  int       `json:"category_id"`
	Image       *string   `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
}

type GetDishDTO struct {
	DishDTO
	Images []string `json:"images"`
}
