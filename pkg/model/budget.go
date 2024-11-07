package model

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

type Budget struct {
	MonthYear string    `json:"month_year,omitempty" gorm:"primaryKey"`
	Amount    int       `json:"amount,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}

// Print preety prints Budget object
func (b *Budget) Print() {
	monthYear := strings.Split(b.MonthYear, "-")
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprint(w, "Year\tMonth\tAmount\n")
	fmt.Fprintf(
		w,
		"%s\t%s\t%d\n",
		monthYear[1],
		monthYear[0],
		b.Amount,
	)
	w.Flush()
}

type Budgets []Budget

// Print preety prints Budget data
func (b Budgets) Print() {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprint(w, "Year\tMonth\tAmount\n")
	for _, budg := range b {
		monthYear := strings.Split(budg.MonthYear, "-")
		fmt.Fprintf(
			w,
			"%s\t%s\t%d\n",
			monthYear[1],
			monthYear[0],
			budg.Amount,
		)
	}
	w.Flush()
}
