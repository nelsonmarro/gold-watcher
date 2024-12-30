package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/nelsonmarro/gold-watcher/internal/models"
)

var (
	errUpdateFailed = errors.New("failed to update the holding")
	errDeleteFailed = errors.New("failed to delete the holding")
)

type HoldingRepository interface {
	Create(holding models.Holding) (*models.Holding, error)
	GetAll() ([]models.Holding, error)
	GetByID(id int64) (*models.Holding, error)
	Update(id int64, holding models.Holding) error
	Delete(id int64) error
}

type holdingRepository struct {
	db *sql.DB
}

func NewHoldingRepository(db *sql.DB) *holdingRepository {
	return &holdingRepository{db: db}
}

func (r *holdingRepository) Create(holding models.Holding) (*models.Holding, error) {
	cmd := `INSERT INTO holdings (amount, purchase_date, purchase_price) VALUES (?, ?, ?);`

	result, err := r.db.Exec(cmd, holding.Amount, holding.PurchaseDate.Unix(), holding.PurchasePrice)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	holding.ID = id
	return &holding, nil
}

func (r *holdingRepository) GetAll() ([]models.Holding, error) {
	query := `SELECT * FROM holdings;`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	holdings := make([]models.Holding, 0)
	unixTime := int64(0)

	for rows.Next() {
		holding := models.Holding{}
		err := rows.Scan(&holding.ID, &holding.Amount, &unixTime, &holding.PurchasePrice)
		if err != nil {
			return nil, err
		}
		holding.PurchaseDate = time.Unix(unixTime, 0)
		holdings = append(holdings, holding)
	}

	return holdings, nil
}

func (r *holdingRepository) GetByID(id int64) (*models.Holding, error) {
	row := r.db.QueryRow(`SELECT * FROM holdings WHERE id = ?;`, id)

	var h models.Holding
	var unixTime int64
	err := row.Scan(&h.ID, &h.Amount, &unixTime, &h.PurchasePrice)
	if err != nil {
		return nil, err
	}

	h.PurchaseDate = time.Unix(unixTime, 0)
	return &h, nil
}

func (r *holdingRepository) Update(id int64, holding models.Holding) error {
	if id == 0 {
		return errors.New("invalid updated id")
	}

	cmd := `UPDATE holdings SET amount = ?, purchase_date = ?, purchase_price = ? WHERE id = ?;`

	res, err := r.db.Exec(cmd, holding.Amount, holding.PurchaseDate.Unix(), holding.PurchasePrice, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errUpdateFailed
	}

	return nil
}

func (r *holdingRepository) Delete(id int64) error {
	if id == 0 {
		return errors.New("invalid deleted id")
	}

	cmd := `DELETE FROM holdings WHERE id = ?;`

	res, err := r.db.Exec(cmd, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errDeleteFailed
	}

	return nil
}
