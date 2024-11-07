package db

import (
	"gorm.io/gorm"

	"github.com/DragoHex/expense-tracker/pkg/model"
)

type DBRepoBudgetImpl struct {
	db *gorm.DB
}

func NewDBRepoBudgetImpl(db *gorm.DB) *DBRepoBudgetImpl {
	return &DBRepoBudgetImpl{db: db}
}

// Create saves an budget entry to DB
func (r *DBRepoBudgetImpl) Create(bud *model.Budget) error {
	return r.db.Create(bud).Error
}

// Read reads an budget entry for DB for the passed monthYear
func (r *DBRepoBudgetImpl) Read(monthYear string) (*model.Budget, error) {
	var bud model.Budget

	err := r.db.Where("month_year = ?", monthYear).First(&bud).Error
	if err != nil {
		return nil, err
	}
	return &bud, nil
}

// Update updates the existing budget entry in DB
func (r *DBRepoBudgetImpl) Update(bud *model.Budget) error {
	return r.db.Save(bud).Error
}

// Delete deletes the budget whose month-year is passed
func (r *DBRepoBudgetImpl) Delete(monthYear string) error {
	return r.db.Where("month_year = ?", monthYear).Delete(&model.Budget{}).Error
}

// List fetches all the budget from the DB
func (r *DBRepoBudgetImpl) List() (model.Budgets, error) {
	var budgets model.Budgets
	err := r.db.Find(&budgets).Error
	if err != nil {
		return nil, err
	}
	return budgets, nil
}
