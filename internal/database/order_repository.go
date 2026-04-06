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

func (r *OrderRepository) GetSlotsByID(id int64) (models.Slot, error) {
	var slot models.Slot
	query := `SELECT id,venue,start_at,end_at,busy FROM slots WHERE id=$1;`
	err := r.db.QueryRow(context.Background(), query, id).Scan(
		&slot.ID,
		&slot.Venue,
		&slot.StartAt,
		&slot.EndAt,
		&slot.Busy,
	)
	if err != nil {
		return models.Slot{}, err
	}
	return slot, nil

}

func (r *OrderRepository) CreateOrder(input models.CreateOrderInput) (models.Order, error) {
	var order models.Order
	query := `INSERT INTO orders (slot_id,customer_name,status) VALUES ($1,$2,'pending')
	returning id,slot_id,customer_name,status,created_at`
	err := r.db.QueryRow(context.Background(), query, input.SlotID, input.CustomerName).Scan(
		&order.ID,
		&order.SlotID,
		&order.CustomerName,
		&order.Status,
		&order.CreatedAt,
	)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}
