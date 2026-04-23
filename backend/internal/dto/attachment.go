package dto

import "time"

type CreateAttachmentDTO struct {
	DishID int    `json:"dish_id"`
	Key    string `json:"key,omitempty"`
	Ext    string `json:"ext,omitempty"`
}

type AttachmentDTO struct {
	ID        int       `json:"id"`
	DishID    int       `json:"dish_id"`
	Key       string    `json:"key"`
	CreatedAt time.Time `json:"created_at"`
}
