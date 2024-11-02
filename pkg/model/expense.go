package model

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type Category int

const (
	_ Category = iota
	Uncategorised
	Groceries
	Transport
	Medical
	Rent
	Entertainment
)

func (c Category) String() string {
	return [...]string{
		"uncategorised",
		"groceries",
		"transport",
		"medical",
		"rent",
		"entertainment",
	}[c-1]
}

func (c Category) EnumIndex() int {
	return int(c)
}

func StringToCatEnum(cat string) int {
	switch cat {
	case "groceries":
		return Groceries.EnumIndex()
	case "transport":
		return Transport.EnumIndex()
	case "medical":
		return Medical.EnumIndex()
	case "rent":
		return Rent.EnumIndex()
	case "entertainment":
		return Entertainment.EnumIndex()
	}
	return Uncategorised.EnumIndex()
}

type Expense struct {
	ID          int       `json:"id,omitempty"          gorm:"primaryKey;autoIncrement"`
	Description string    `json:"description,omitempty"`
	Amount      int       `json:"amount,omitempty"`
	Category    int       `json:"category,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"  gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"  gorm:"autoUpdateTime"`
}

func (e Expense) Print() {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprint(w, "ID\tDate\tDescription\tAmount\tCategory\n")
	fmt.Fprintf(
		w,
		"%d\t%v\t%s\t%d\t%s\n",
		e.ID,
		e.CreatedAt.Format("2006-01-02"),
		e.Description,
		e.Amount,
		Category(e.Category).String(),
	)
	w.Flush()
}

type Expenses []Expense

func (e Expenses) Print() {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprint(w, "ID\tDate\tDescription\tAmount\tCategory\n")
	for _, exp := range e {
		fmt.Fprintf(
			w,
			"%d\t%v\t%s\t%d\t%s\n",
			exp.ID,
			exp.CreatedAt.Format("2006-01-02"),
			exp.Description,
			exp.Amount,
			Category(exp.Category).String(),
		)
	}
	w.Flush()
}

// Summary returns total expenditure for the current year
// if a valid month is passed then expenditure in that month is shown
func (e Expenses) Summary(m int, y int) {
	var total int
	resp := "Total expenses"
	month := time.Month(m)

	fmt.Println(y)
	switch {
	case month >= time.January && month <= time.December:
		for _, exp := range e {
			if exp.CreatedAt.Month() == month && exp.CreatedAt.Year() == y {
				total = total + exp.Amount
			}
		}
		resp = fmt.Sprintf("%s for %s: %d", resp, month.String(), total)
	default:
		for _, exp := range e {
			if exp.CreatedAt.Year() == y {
				total = total + exp.Amount
			}
		}
		resp = fmt.Sprintf("%s: %d", resp, total)
	}

	fmt.Println(resp)
}
