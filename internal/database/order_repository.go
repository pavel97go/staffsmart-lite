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

func (r *OrderRepository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	query := `SELECT id, slot_id,customer_name,status,created_at FROM orders
	ORDER BY id DESC`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var order models.Order
		err := rows.Scan(
			&order.ID,
			&order.SlotID,
			&order.CustomerName,
			&order.Status,
			&order.CreatedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) GetOrderByID(id int64) (models.Order, error) {
	var order models.Order
	query := `SELECT id, slot_id,customer_name,status,created_at FROM orders
WHERE id = $1`
	err := r.db.QueryRow(context.Background(), query, id).Scan(
		&order.ID,
		&order.SlotID,
		&order.CustomerName,
		&order.Status,
		&order.CreatedAt)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}
