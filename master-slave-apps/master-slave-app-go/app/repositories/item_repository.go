package repositories

import (
	"database/sql"
	"time"

	"apps/config"
	"apps/models"
)

type ItemRepository struct {
	masterDB *sql.DB
	slaveDB  *sql.DB
}

func NewItemRepository(config *config.Config) (*ItemRepository, error) {
	return &ItemRepository{
		masterDB: config.MasterDB,
		slaveDB:  config.SlaveDB,
	}, nil
}

func (r *ItemRepository) Create(item *models.Item) error {
	query := "INSERT INTO items (title, description, qty, price, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	now := time.Now()
	result, err := r.masterDB.Exec(query, item.Title, item.Description, item.Quantity, item.Price, now, now)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	item.CreatedAt = models.CustomTime{now}
	item.ID = int(id)
	return nil
}

func (r *ItemRepository) FindAll() ([]models.Item, error) {
	rows, err := r.slaveDB.Query("SELECT id, title, description, qty, price, created_at, updated_at FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		var tc, tu time.Time
		err := rows.Scan(&item.ID, &item.Title, &item.Description, &item.Quantity, &item.Price, &tc, &tu)
		if err != nil {
			return nil, err
		}

		item.CreatedAt = models.CustomTime{tc}
		item.UpdatedAt = &models.CustomTime{tu}

		items = append(items, item)
	}
	return items, nil
}

func (r *ItemRepository) FindByID(id string) (models.Item, error) {
	var item models.Item
	var tc, tu time.Time
	err := r.slaveDB.QueryRow("SELECT id, title, description, qty, price, created_at, updated_at FROM items WHERE id = ?", id).
		Scan(&item.ID, &item.Title, &item.Description, &item.Quantity, &item.Price, &tc, &tu)
		item.CreatedAt = models.CustomTime{tc}
		item.UpdatedAt = &models.CustomTime{tu}
	return item, err
}

func (r *ItemRepository) Update(id string, item *models.Item) (int64, error) {
	query := "UPDATE items SET title = ?, description = ?, qty = ?, price = ?, updated_at = ? WHERE id = ?"
	result, err := r.masterDB.Exec(query, item.Title, item.Description, item.Quantity, item.Price, time.Now(), id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (r *ItemRepository) Delete(id string) (int64, error) {
	result, err := r.masterDB.Exec("DELETE FROM items WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}