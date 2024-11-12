package tracker

import (
	"github.com/DragoHex/expense-tracker/pkg/model"
	expense "github.com/DragoHex/expense-tracker/pkg/model"
)

type ExpenseTracker interface {
	CreateExpense(des string, amount, cat int) (*expense.Expense, error)
	GetExpense(id int) (*expense.Expense, error)
	UpdateExpense(id int, des string, amount int, cat int) error
	DeleteExpense(id int) error
	ListExpense() ([]expense.Expense, error)
}

type ExpenseTrackerImpl struct {
	repo ExpenseRepository
}

// NewExpenseTrackerImpl returns an instance of ExpenseRepository
func NewExpenseTrackerImpl(repo ExpenseRepository) *ExpenseTrackerImpl {
	return &ExpenseTrackerImpl{repo: repo}
}

// CreateExpense adds a new expense entry
func (s *ExpenseTrackerImpl) CreateExpense(des string, amount, cat int) (*expense.Expense, error) {
	exp := &expense.Expense{Description: des, Amount: amount, Category: model.Category(cat)}

	err := s.repo.Create(exp)
	if err != nil {
		return nil, err
	}

	return exp, nil
}

// GetExpense gets expense as per the passed id
func (s *ExpenseTrackerImpl) GetExpense(id int) (*expense.Expense, error) {
	return s.repo.Read(id)
}

// UpdateExpense updates an existing expenses
func (s *ExpenseTrackerImpl) UpdateExpense(id int, des string, amount int, cat int) error {
	exp, err := s.repo.Read(id)
	if err != nil {
		return err
	}

	exp.Description = des
	exp.Amount = amount
	exp.Category = model.Category(cat)

	return s.repo.Update(exp)
}

// DeleteExpense deletes the expense whose id is passed
func (s *ExpenseTrackerImpl) DeleteExpense(id int) error {
	return s.repo.Delete(id)
}

// ListExpense fetches all the expenses
func (s *ExpenseTrackerImpl) ListExpense() (expense.Expenses, error) {
	return s.repo.List()
}

// ListMonthyExpense fetches expenses of the particular month
func (s *ExpenseTrackerImpl) ListFilteredExpense(m, y int) (expense.Expenses, error) {
	return s.repo.ListFiltered(m, y)
}
