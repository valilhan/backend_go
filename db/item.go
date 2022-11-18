package db

import (
	"context"
	"database/sql"

	"gitlab.com/idoko/HyperSkill/models"
)

func (db Database) AddBalance(item *models.Balance) error {
	ctx := context.Background()
	tx, err := db.Conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var CreatedAt string
	query := `INSERT INTO main_balance (user_id, balance) VALUES ($1, $2) RETURNING created_at;`
	if err := db.Conn.QueryRowContext(ctx, query, item.UserId, item.Balance).Scan(&CreatedAt); err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil

}

func (db Database) GetBalanceById(get *models.GetBalance) (models.Balance, error) {
	item := models.Balance{} //Empty class
	query := `SELECT user_id, balance FROM main_balance WHERE user_id = $1;`
	row := db.Conn.QueryRow(query, get.UserId)
	switch err := row.Scan(&item.UserId, &item.Balance); err {
	case sql.ErrNoRows:
		return item, ErrNoMatch
	default:
		return item, err
	}
}

func (db Database) AddReserveBalance(item *models.ReserveBalance) error {
	var CreatedAt string
	ctx := context.Background()

	tx, err := db.Conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	query := `INSERT INTO reserve_balance (user_id, service_id, order_id, price) 
	VALUES ($1, $2, $3, $4 ) RETURNING created_at;`

	if err := tx.QueryRowContext(ctx, query, item.UserId, item.ServiceId, item.OrderId, item.Price).Scan(&CreatedAt); err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (db Database) AddRevenue(item *models.Revenue) error {
	ctx := context.Background()
	// Get a Tx for making transaction requests.

	tx, err := db.Conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()
	query := `SELECT price FROM reserve_balance WHERE user_id = $1 and service_id = $2 and order_id = $3;`
	var price float64

	if err := tx.QueryRowContext(ctx, query, item.UserId, item.ServiceId, item.OrderId).Scan(&price); err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()
		return err
	}
	query = `UPDATE main_balance SET balance = balance - $1;`
	if _, err := tx.ExecContext(ctx, query, price); err != nil {
		tx.Rollback()
		return err
	}

	query = `INSERT INTO revenue (user_id, service_id, order_id, price) VALUES ($1, $2, $3, $4 ) RETURNING created_at;`

	var CreatedAt string
	if err := db.Conn.QueryRowContext(ctx, query, item.UserId, item.ServiceId, item.OrderId, item.Price).Scan(&CreatedAt); err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
