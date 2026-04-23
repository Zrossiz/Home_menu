package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Zrossiz/Home_menu/backend/internal/dto"
)

type AttachmentPostgresStorage struct {
	conn *sql.DB
}

func NewAttchmentPostgresStorage(conn *sql.DB) *AttachmentPostgresStorage {
	return &AttachmentPostgresStorage{conn: conn}
}

func (a *AttachmentPostgresStorage) Create(payload dto.CreateAttachmentDTO) error {
	query := `
		INSERT INTO 
			attachments (dish_id, key)
		VALUES
			($1, $2)
	`

	_, err := a.conn.Exec(query, payload.DishID, payload.Key)
	if err != nil {
		return fmt.Errorf("error create attachment: %v", err)
	}

	return nil
}

func (a *AttachmentPostgresStorage) GetAllByDish(dishID int) ([]dto.AttachmentDTO, error) {
	query := `
		SELECT 
			id, 
			dish_id, 
			key, 
			created_at 
		FROM 
			attachments 
		WHERE 
			dish_id = $1`

	rows, err := a.conn.Query(query, dishID)
	if err != nil {
		return nil, fmt.Errorf("error get attachments rows: %v", err)
	}
	defer rows.Close()

	var attachments []dto.AttachmentDTO
	for rows.Next() {
		var attachment dto.AttachmentDTO

		err = rows.Scan(
			&attachment.ID,
			&attachment.DishID,
			&attachment.Key,
			&attachment.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scan attachment row: %v", err)
		}

		attachments = append(attachments, attachment)
	}

	return attachments, nil
}
