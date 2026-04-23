package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	Category   CategoryPostgresStorage
	Dish       DishPostgresStorage
	Attachment AttachmentPostgresStorage
}

func New(conn *sql.DB) *PostgresStorage {
	var store PostgresStorage

	store.Category = *NewCategoryPostgresStorage(conn)
	store.Dish = *NewDishStorage(conn)
	store.Attachment = *NewAttchmentPostgresStorage(conn)

	return &store
}

func Connect(DBURI string) (*sql.DB, error) {
	conn, err := sql.Open("postgres", DBURI)
	if err != nil {
		return nil, fmt.Errorf("error connect to DB: %v", err)
	}

	return conn, nil
}
