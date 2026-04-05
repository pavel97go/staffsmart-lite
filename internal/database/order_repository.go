package database

import (
	"context"
	"staffsmart-lite/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) GetAllSlots() ([]models.Slot, error) {
	var slots []models.Slot
	query := `
	SELECT id,venue,start_at,end_at,busy FROM slots
	ORDER BY id desc`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var slot models.Slot
		err := rows.Scan(&slot.ID,
			&slot.Venue,
			&slot.StartAt,
			&slot.EndAt,
			&slot.Busy,
		)
		if err != nil {
			return nil, err
		}
		slots = append(slots, slot)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return slots, nil
}
