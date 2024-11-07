package cmd

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/DragoHex/expense-tracker/pkg/model"
	"github.com/spf13/cobra"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "A command to get expense summary",
	Long: `A command to get expense summary

If no month/year is passed it will give summary for the current month.
If only year is passed it will give summary for the whole year`,
	Run: func(cmd *cobra.Command, args []string) {
		month, err := cmd.Flags().GetInt("month")
		if err != nil {
			fmt.Printf("error in getting the flag: %s\n", err)
			return
		}
		year, err := cmd.Flags().GetInt("year")
		if err != nil {
			fmt.Printf("error in getting the flag: %s\n", err)
			return
		}

		if month > 12 || month < 0 {
			fmt.Println("invalid month enter")
			fmt.Println("valid values are [1, 12]")
			return
		}

		cat, err := cmd.Flags().GetString("category")
		if err != nil {
			fmt.Printf("error in getting the flag value: %s", err)
			return
		}

		var exps model.Expenses

		// Fetch expense for a particular month
		if !cmd.Flags().Changed("year") {
			exps, err = ExpenseTracker.ListFilteredExpense(month, year)
			if err != nil {
				fmt.Printf("error in fetching the expenses: %s\n", err)
				return
			}
		}

		// Fetch yearly expenses
		if cmd.Flags().Changed("year") && !cmd.Flags().Changed("month") {
			exps, err = ExpenseTracker.ListFilteredExpense(0, year)
			if err != nil {
				fmt.Printf("error in fetching the expenses: %s\n", err)
				return
			}
		}

		if cat != "" {
			cats := strings.Split(cat, ",")
			var categoricalExpenses model.Expenses
			for _, exp := range exps {
				if slices.Contains(cats, exp.Category.String()) {
					categoricalExpenses = append(categoricalExpenses, exp)
				}
			}
			exps = categoricalExpenses
		}
		exps.Summary(month, year)
	},
}

func init() {
	t := time.Now()
	summaryCmd.Flags().IntP("month", "m", int(t.Month()), "month")
	summaryCmd.Flags().IntP("year", "y", t.Year(), "year, by default it's current year")
	summaryCmd.Flags().StringP("category", "c", "", `expense categories, multiple categories can be passed comma separted
supported categries are: groceries, transport, medical, rent, entertainment`)
	rootCmd.AddCommand(summaryCmd)
}
