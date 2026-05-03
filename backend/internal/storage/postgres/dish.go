package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Zrossiz/Home_menu/backend/internal/apperrors"
	"github.com/Zrossiz/Home_menu/backend/internal/dto"
)

type DishPostgresStorage struct {
	conn *sql.DB
}

func NewDishStorage(conn *sql.DB) *DishPostgresStorage {
	return &DishPostgresStorage{conn: conn}
}

func (d *DishPostgresStorage) Create(payload dto.CreateDishDTO) error {
	query := `
		INSERT INTO 
			dishes (category_id, name, recipe, description, time_to_cook) 
		VALUES 
			($1, $2, $3, $4, $5)
	`

	_, err := d.conn.Exec(
		query,
		payload.CategoryID,
		payload.Name,
		payload.Recipe,
		payload.Description,
		payload.TimeToCook,
	)
	if err != nil {
		return fmt.Errorf("error insert dished: %v", err)
	}

	return nil
}

func (d *DishPostgresStorage) GetAllByCategory(categoryID int) ([]dto.DishDTO, error) {
	query := `
		SELECT 
			d.id, 
			d.category_id, 
			d.name, 
			d.time_to_cook, 
			d.recipe, 
			d.description,
			a.key,
			d.created_at 
		FROM 
			dishes d
		LEFT JOIN LATERAL (
			SELECT *
			FROM attachments a
			WHERE d.id = a.dish_id
			ORDER BY a.id
			LIMIT 1
		) a ON true
		WHERE
			category_id = $1
	`

	rows, err := d.conn.Query(query, categoryID)
	if err != nil {
		return nil, fmt.Errorf("error get dish rows: %v", err)
	}
	defer rows.Close()

	var dishes []dto.DishDTO
	for rows.Next() {
		var dish dto.DishDTO

		err := rows.Scan(
			&dish.ID,
			&dish.CategoryID,
			&dish.Name,
			&dish.TimeToCook,
			&dish.Recipe,
			&dish.Description,
			&dish.Image,
			&dish.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scan dish row: %v", err)
		}

		dishes = append(dishes, dish)
	}

	return dishes, nil
}

func (d *DishPostgresStorage) Update(dishID int, payload dto.UpdateDishDTO) error {
	query := `
		UPDATE 
			dishes 
		SET 
			name = $1, 
			description = $2, 
			recipe = $3, 
			time_to_cook = $4
		WHERE
			id = $5
	`

	res, err := d.conn.Exec(
		query,
		payload.Name,
		payload.Description,
		payload.Recipe,
		payload.TimeToCook,
		dishID,
	)
	if err != nil {
		return fmt.Errorf("error update dish: %v", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error get dish affected rows: %v", err)
	}
	if count == 0 {
		return apperrors.ErrNotFound
	}

	return nil
}

func (d *DishPostgresStorage) Delete(dishID int) error {
	query := `
		DELETE FROM dishes WHERE id = $1 
	`

	res, err := d.conn.Exec(query, dishID)
	if err != nil {
		return fmt.Errorf("error delete dish: %v", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error get dish affected rows: %v", err)
	}
	if count == 0 {
		return apperrors.ErrNotFound
	}

	return nil
}

func (d *DishPostgresStorage) GetOne(dishId int) (*dto.DishDTO, error) {
	query := `
		SELECT 
			id, category_id, name, description, recipe, time_to_cook, created_at
		FROM 
			dishes
		WHERE
			id = $1
	`

	row := d.conn.QueryRow(query, dishId)

	var dish dto.DishDTO
	err := row.Scan(
		&dish.ID,
		&dish.CategoryID,
		&dish.Name,
		&dish.Description,
		&dish.Recipe,
		&dish.TimeToCook,
		&dish.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrNotFound
		}
		return nil, fmt.Errorf("error get dish by id: %v", err)
	}

	return &dish, nil
}

func (d *DishPostgresStorage) Find(search string) ([]dto.DishDTO, error) {
	query := `
		SELECT 
			d.id,
			d.category_id,
			d.name,
			d.time_to_cook,
			d.description,
			a.key,
			d.created_at
		FROM dishes d
		LEFT JOIN LATERAL (
			SELECT *
			FROM attachments a
			WHERE d.id = a.dish_id
			ORDER BY a.id
			LIMIT 1
		) a ON true
		WHERE name ILIKE '%' || $1 || '%';
	`

	rows, err := d.conn.Query(query, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []dto.DishDTO
	for rows.Next() {
		var dish dto.DishDTO
		err = rows.Scan(
			&dish.ID,
			&dish.CategoryID,
			&dish.Name,
			&dish.TimeToCook,
			&dish.Description,
			&dish.Image,
			&dish.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, dish)
	}

	return result, nil
}
