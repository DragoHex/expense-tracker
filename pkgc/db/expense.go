package db

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

// Print preety print expense object in a table
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

// Print preety prints the Expense data in a table
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
func (e Expenses) Summary(m, y int) {
	resp := "Total expenses"
	month := time.Month(m)
	total := e.ConditionalTotal(m, y)

	switch {
	case month >= time.January && month <= time.December:
		resp = fmt.Sprintf("%s for %s: %d", resp, month.String(), total)
	default:
		resp = fmt.Sprintf("%s: %d", resp, total)
	}

	fmt.Println(resp)
}

// ConditionalTotal sums up expenses for the given time period
func (e Expenses) ConditionalTotal(m, y int) int {
	var total int
	month := time.Month(m)

	switch {
	case month >= time.January && month <= time.December:
		for _, exp := range e {
			if exp.CreatedAt.Month() == month && exp.CreatedAt.Year() == y {
				total = total + exp.Amount
			}
		}
	default:
		for _, exp := range e {
			if exp.CreatedAt.Year() == y {
				total = total + exp.Amount
			}
		}
	}
	return total
}

// Total sums up all the expenses
func (e Expenses) Total() int {
	var total int
	for _, exp := range e {
		total = total + exp.Amount
	}
	return total
}
