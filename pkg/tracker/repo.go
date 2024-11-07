package tracker

import (
	"github.com/DragoHex/expense-tracker/pkg/model"
)

type ExpenseRepository interface {
	Create(model *model.Expense) error
	Read(id int) (*model.Expense, error)
	Update(model *model.Expense) error
	Delete(id int) error
	List() (model.Expenses, error)
	ListFiltered(m, y int) (model.Expenses, error)
}

type BudgetRepository interface {
	Create(model *model.Budget) error
	Read(monthYear string) (*model.Budget, error)
	Update(model *model.Budget) error
	Delete(monthYear string) error
	List() (model.Budgets, error)
}
