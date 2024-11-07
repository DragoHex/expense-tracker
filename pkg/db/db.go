package db

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/DragoHex/expense-tracker/pkg/model"
)

type DBRepoImpl struct {
	db *gorm.DB
}

func NewDBRepoImpl(db *gorm.DB) *DBRepoImpl {
	return &DBRepoImpl{db: db}
}

// Create saves an model entry to DB
func (r *DBRepoImpl) Create(exp *model.Expense) error {
	return r.db.Create(exp).Error
}

// Read reads an model entry for DB for the passed id
func (r *DBRepoImpl) Read(id int) (*model.Expense, error) {
	var exp model.Expense

	err := r.db.First(&exp, id).Error
	if err != nil {
		return nil, err
	}
	return &exp, nil
}

// Update updates the existing model entry in DB
func (r *DBRepoImpl) Update(exp *model.Expense) error {
	return r.db.Save(exp).Error
}

// Delete deletes the model whose id is passed
func (r *DBRepoImpl) Delete(id int) error {
	return r.db.Delete(&model.Expense{}, id).Error
}

// List fetches all the models from the DB
func (r *DBRepoImpl) List() (model.Expenses, error) {
	var exps model.Expenses
	err := r.db.Find(&exps).Error
	if err != nil {
		return nil, err
	}
	return exps, nil
}

// ListFiltered fetches all the models from the DB
func (r *DBRepoImpl) ListFiltered(m, y int) (model.Expenses, error) {
	var exps model.Expenses

	if m == 0 {
		err := r.db.Where("strftime('%Y', created_at) = ?", fmt.Sprintf("%d", y)).
			Find(&exps).
			Error
		if err != nil {
			return nil, err
		}
		return exps, nil
	}

	err := r.db.Where("strftime('%m', created_at) = ? AND strftime('%Y', created_at) = ?", fmt.Sprintf("%02d", m), fmt.Sprintf("%d", y)).
		Find(&exps).
		Error
	if err != nil {
		return nil, err
	}
	return exps, nil
}
