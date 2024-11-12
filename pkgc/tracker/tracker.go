package tracker

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/DragoHex/expense-tracker/pkgc/db"
)

type ExpenseTracker interface {
	CreateExpense(des string, amount, cat int) (db.Expense, error)
	GetExpense(id int) (db.Expense, error)
	UpdateExpense(id int, des string, amount int, cat int) error
	DeleteExpense(id int) error
	ListExpense() (db.Expense, error)
}

type ExpenseTrackerImpl struct {
	dbObj *sql.DB
}

// NewExpenseTrackerImpl returns an instance of ExpenseRepository
func NewExpenseTrackerImpl(dbObj *sql.DB) *ExpenseTrackerImpl {
	return &ExpenseTrackerImpl{dbObj: dbObj}
}

// CreateExpense adds a new expense entry
func (s *ExpenseTrackerImpl) CreateExpense(des string, amount int, cat string) (db.Expense, error) {
	txn, err := s.dbObj.Begin()
	if err != nil {
		return db.Expense{}, err
	}

	dbQueries := db.New(txn)
	ctx := context.Background()

	exp, err := dbQueries.CreateExpense(ctx, db.CreateExpenseParams{
		Description: des,
		Amount:      amount,
		Category:    db.StringToCatEnum(cat),
	})
	if err != nil {
		txn.Rollback()
		return db.Expense{}, err
	}

	err = txn.Commit()
	if err != nil {
		return db.Expense{}, err
	}

	return exp, nil
}

// GetExpense gets expense as per the passed id
func (s *ExpenseTrackerImpl) GetExpense(id int) (db.Expense, error) {
	txn, err := s.dbObj.Begin()
	if err != nil {
		return db.Expense{}, err
	}

	dbQueries := db.New(txn)
	ctx := context.Background()

	exp, err := dbQueries.GetExpense(ctx, id)

	err = txn.Commit()
	if err != nil {
		return db.Expense{}, err
	}

	return exp, nil
}

// UpdateExpense updates an existing expenses
func (s *ExpenseTrackerImpl) UpdateExpense(id int, des string, amount int, cat int) error {
	txn, err := s.dbObj.Begin()
	if err != nil {
		return err
	}

	dbQueries := db.New(txn)
	ctx := context.Background()

	exp, err := dbQueries.GetExpense(ctx, id)
	if err != nil {
		txn.Rollback()
		return err
	}

	if des == "" {
		des = exp.Description
	}
	if cat == 0 {
		cat = exp.Category
	}
	if amount == 0 {
		amount = exp.Amount
	}

	err = dbQueries.UpdateExpense(
		ctx,
		db.UpdateExpenseParams{
			ID:          id,
			Description: des,
			Amount:      amount,
			Category:    cat,
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

// DeleteExpense deletes the expense whose id is passed
func (s *ExpenseTrackerImpl) DeleteExpense(id int) error {
	txn, err := s.dbObj.Begin()
	if err != nil {
		return err
	}

	dbQueries := db.New(txn)
	ctx := context.Background()

	err = dbQueries.DeleteExpense(ctx, id)
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

// ListExpense fetches all the expenses
func (s *ExpenseTrackerImpl) ListExpense() (db.Expenses, error) {
	txn, err := s.dbObj.Begin()
	if err != nil {
		return nil, err
	}

	dbQueries := db.New(txn)
	ctx := context.Background()

	exps, err := dbQueries.ListExpense(ctx)
	if err != nil {
		txn.Rollback()
		return nil, err
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	return exps, nil
}

// ListMonthyExpense fetches expenses of the particular month
func (s *ExpenseTrackerImpl) ListFilteredExpense(m, y int) (db.Expenses, error) {
	txn, err := s.dbObj.Begin()
	if err != nil {
		return nil, err
	}

	dbQueries := db.New(txn)
	ctx := context.Background()

	exps, err := dbQueries.ListExpense(ctx)
	exps, err = dbQueries.ListFilteredExpense(
		ctx,
		db.ListFilteredExpenseParams{
			Column1: fmt.Sprintf("%d", y),
			Column2: fmt.Sprintf("%02d", m),
		},
	)
	if err != nil {
		txn.Rollback()
		return nil, err
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	return exps, nil
}
