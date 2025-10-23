package database

import (
	"database/sql"

	"github.com/google/uuid"
)

// Category representa uma categoria no banco de dados.
type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

// NewCategory cria uma nova instância de Category.
func NewCategory(db *sql.DB, id, name, description string) *Category {
	return &Category{db: db}
}

// Create insere uma nova categoria no banco de dados.
func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES (?, ?, ?)",
		id, name, description)

	if err != nil {
		return Category{}, err
	}

	return Category{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}
