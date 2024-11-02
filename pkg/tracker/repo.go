package tracker

import (
	expense "github.com/DragoHex/expense-tracker/pkg/model"
)

type ExpenseRepository interface {
	Create(expense *expense.Expense) error
	Read(id int) (*expense.Expense, error)
	Update(expense *expense.Expense) error
	Delete(id int) error
	List() (expense.Expenses, error)
}
