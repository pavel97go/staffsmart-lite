package database

import (
	"context"
	"database/sql"
	"staffsmart-lite/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SlotRepository struct {
	db *pgxpool.Pool
}

func NewSlotRepository(db *pgxpool.Pool) *SlotRepository {
	return &SlotRepository{db: db}
}

func (r *SlotRepository) CreateSlot(slot models.Slot) (models.Slot, error) {
	var createdSlot models.Slot
	query := `INSERT INTO slots (venue, start_at, end_at, busy)
	VALUES ($1, $2, $3, $4)
	RETURNING id, venue, start_at, end_at, busy`
	err := r.db.QueryRow(context.Background(), query, slot.Venue, slot.StartAt, slot.EndAt, slot.Busy).Scan(
		&createdSlot.ID,
		&createdSlot.Venue,
		&createdSlot.StartAt,
		&createdSlot.EndAt,
		&createdSlot.Busy)
	if err != nil {
		return models.Slot{}, err
	}
	return createdSlot, nil

}

func (r *SlotRepository) UpdateSlot(id int64, slot models.Slot) (models.Slot, error) {
	var updatedSlot models.Slot
	query := `
		UPDATE slots
		SET venue = $1, start_at = $2, end_at = $3, busy = $4
		WHERE id = $5
		RETURNING id, venue, start_at, end_at, busy
	`

	err := r.db.QueryRow(context.Background(), query, slot.Venue, slot.StartAt, slot.EndAt, slot.Busy, id).Scan(
		&updatedSlot.ID,
		&updatedSlot.Venue,
		&updatedSlot.StartAt,
		&updatedSlot.EndAt,
		&updatedSlot.Busy,
	)

	if err != nil {
		return models.Slot{}, err
	}

	return updatedSlot, nil
}

func (r *SlotRepository) DeleteSlot(id int64) error {
	query := `DELETE FROM slots WHERE id = $1`
	tag, err := r.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *SlotRepository) GetAllSlots() ([]models.Slot, error) {
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

func (r *SlotRepository) GetSlotByID(id int64) (models.Slot, error) {
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
