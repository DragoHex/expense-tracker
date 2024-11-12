package tracker

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/DragoHex/expense-tracker/pkgc/db"
	"github.com/DragoHex/expense-tracker/pkgc/utils"
)

type Budgeteer interface {
	CreateBudget(m, y, amount int) (db.Budget, error)
	GetBudget(m, y int) (db.Budget, error)
	UdateBudget(m, y, amount int) error
	DeleteBudget(m, y int) error
	ListBudget() (db.Budgets, error)
}

type BudgetRepoImpl struct {
	dbObj *sql.DB
}

// NewBudgetRepoImpl returns an instance of ExpenseRepository
func NewBudgetRepoImpl(dbObj *sql.DB) *BudgetRepoImpl {
	return &BudgetRepoImpl{dbObj: dbObj}
}

// CreateBudget create a new budget entry
func (b *BudgetRepoImpl) CreateBudget(m, y, amount int) (db.Budget, error) {
	if m < 1 && m > 12 {
		return db.Budget{}, fmt.Errorf("invalid month passed: %d", m)
	}

	txn, err := b.dbObj.Begin()
	if err != nil {
		return db.Budget{}, err
	}

	dbQueries := db.New(txn)
	ctx := context.Background()

	budg, err := dbQueries.CreateBudget(
		ctx,
		db.CreateBudgetParams{
			MonthYear: utils.ConvertTimetToString(m, y),
			Amount:    amount,
		},
	)
	if err != nil {
		txn.Rollback()
		return db.Budget{}, err
	}

	err = txn.Commit()
	if err != nil {
		return db.Budget{}, err
	}

	return budg, nil
}

// GetBudget fetched budget for the passed month-year
func (b *BudgetRepoImpl) GetBudget(m, y int) (db.Budget, error) {
	txn, err := b.dbObj.Begin()
	if err != nil {
		return db.Budget{}, err
	}

	dbQueries := db.New(txn)
	ctx := context.Background()

	budg, err := dbQueries.GetBudget(ctx, utils.ConvertTimetToString(m, y))

	err = txn.Commit()
	if err != nil {
		return db.Budget{}, err
	}

	return budg, nil
}

// UpdateBudget updates an existing budget entry
func (b *BudgetRepoImpl) UpdateBudget(m, y, amount int) error {
	monthYear := utils.ConvertTimetToString(m, y)
	txn, err := b.dbObj.Begin()
	if err != nil {
		return err
	}

	dbQueries := db.New(txn)
	ctx := context.Background()

	budg, err := dbQueries.GetBudget(ctx, utils.ConvertTimetToString(m, y))

	budg.MonthYear = monthYear
	budg.Amount = amount

	err = dbQueries.UpdateBudget(
		ctx,
		db.UpdateBudgetParams{
			MonthYear: utils.ConvertTimetToString(m, y),
			Amount:    amount,
		},
	)
	if err != nil {
		txn.Rollback()
		return err
	}

	err = txn.Commit()
	if err != nil {
		return err
	}

	return nil
}

// DeleteBudget deletes a budget entry
func (b *BudgetRepoImpl) DeleteBudget(m, y int) error {
	txn, err := b.dbObj.Begin()
	if err != nil {
		return err
	}

	dbQueries := db.New(txn)
	ctx := context.Background()

	err = dbQueries.DeleteBudget(ctx, utils.ConvertTimetToString(m, y))
	if err != nil {
		txn.Rollback()
		return err
	}

	err = txn.Commit()
	if err != nil {
		return err
	}

	return nil
}

// ListBudget lists all the budget entries
func (b *BudgetRepoImpl) ListBudget() (db.Budgets, error) {
	txn, err := b.dbObj.Begin()
	if err != nil {
		return nil, err
	}

	dbQueries := db.New(txn)
	ctx := context.Background()

	budgs, err := dbQueries.ListBudget(ctx)
	if err != nil {
		txn.Rollback()
		return nil, err
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	return budgs, nil
}
