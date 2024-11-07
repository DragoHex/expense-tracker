package tracker

import (
	"fmt"

	"github.com/DragoHex/expense-tracker/pkg/model"
	"github.com/DragoHex/expense-tracker/pkg/utils"
)

type Budgeteer interface {
	CreateBudget(m, y, amount int) (*model.Budget, error)
	GetBudget(m, y int) (*model.Budget, error)
	UdateBudget(m, y, amount int) error
	DeleteBudget(m, y int) error
	ListBudget() (model.Budgets, error)
}

type BudgetRepoImpl struct {
	repo BudgetRepository
}

// NewBudgetRepoImpl returns an instance of ExpenseRepository
func NewBudgetRepoImpl(repo BudgetRepository) *BudgetRepoImpl {
	return &BudgetRepoImpl{repo: repo}
}

// CreateBudget create a new budget entry
func (b *BudgetRepoImpl) CreateBudget(m, y, amount int) (*model.Budget, error) {
	if m < 1 && m > 12 {
		return nil, fmt.Errorf("invalid month passed: %d", m)
	}

	budg := &model.Budget{
		MonthYear: utils.ConvertTimetToString(m, y),
		Amount:    amount,
	}

	return budg, b.repo.Create(budg)
}

// GetBudget fetched budget for the passed month-year
func (b *BudgetRepoImpl) GetBudget(m, y int) (*model.Budget, error) {
	return b.repo.Read(utils.ConvertTimetToString(m, y))
}

// UpdateBudget updates an existing budget entry
func (b *BudgetRepoImpl) UpdateBudget(m, y, amount int) error {
	monthYear := utils.ConvertTimetToString(m, y)
	budg, err := b.repo.Read(monthYear)
	if err != nil {
		return err
	}

	budg.MonthYear = monthYear
	budg.Amount = amount

	return b.repo.Update(budg)
}

// DeleteBudget deletes a budget entry
func (b *BudgetRepoImpl) DeleteBudget(m, y int) error {
	return b.repo.Delete(utils.ConvertTimetToString(m, y))
}

// ListBudget lists all the budget entries
func (b *BudgetRepoImpl) ListBudget() (model.Budgets, error) {
	return b.repo.List()
}
