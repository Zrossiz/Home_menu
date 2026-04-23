package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Zrossiz/Home_menu/backend/internal/apperrors"
	"github.com/Zrossiz/Home_menu/backend/internal/dto"
)

type CategoryPostgresStorage struct {
	conn *sql.DB
}

func NewCategoryPostgresStorage(conn *sql.DB) *CategoryPostgresStorage {
	return &CategoryPostgresStorage{conn: conn}
}

func (c *CategoryPostgresStorage) Create(payload dto.CreateCategoryDTO) error {
	query := `INSERT INTO categories (name) VALUES($1)`

	_, err := c.conn.Exec(query, payload.Name)
	if err != nil {
		return fmt.Errorf("error insert into categories: %v", err)
	}

	return nil
}

func (c *CategoryPostgresStorage) Update(categoryID int, payload dto.UpdateCategoryDTO) error {
	query := `
		UPDATE
			categories
		SET 
			name = $1
		WHERE 
			id = $2
	`

	res, err := c.conn.Exec(query, payload.Name, categoryID)
	if err != nil {
		return fmt.Errorf("update category error: %v", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error get category affected rows")
	}
	if count == 0 {
		return apperrors.ErrNotFound
	}

	return nil
}

func (c *CategoryPostgresStorage) Delete(categoryID int) error {
	query := `
		DELETE FROM 
			categories 
		WHERE 
			id = $1
	`

	res, err := c.conn.Exec(query, categoryID)
	if err != nil {
		return fmt.Errorf("error delete category: %v", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error get category affected rows: %v", err)
	}
	if count == 0 {
		return apperrors.ErrNotFound
	}

	return nil
}

func (c *CategoryPostgresStorage) GetAll() ([]dto.CategoryDTO, error) {
	query := `
		SELECT 
			id, name, created_at
		FROM 
			categories
	`

	rows, err := c.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error get categories: %v", err)
	}
	defer rows.Close()

	var categories []dto.CategoryDTO
	for rows.Next() {
		var category dto.CategoryDTO
		err = rows.Scan(&category.ID, &category.Name, &category.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scan categories rows: %v", err)
		}
		categories = append(categories, category)
	}

	return categories, nil
}
