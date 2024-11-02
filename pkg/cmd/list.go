package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/DragoHex/expense-tracker/pkg/model"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A command to list expenses",
	Long:  `A command to list expenses`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			fmt.Printf("error in getting the flag value: %s\n", err)
			return
		}
		month, err := cmd.Flags().GetInt("month")
		if err != nil {
			fmt.Printf("error in getting the flag value: %s", err)
			return
		}
		year, err := cmd.Flags().GetInt("year")
		if err != nil {
			fmt.Printf("error in getting the flag value: %s", err)
			return
		}
		cat, err := cmd.Flags().GetString("category")
		if err != nil {
			fmt.Printf("error in getting the flag value: %s", err)
			return
		}

		if id != 0 {
			exp, err := ExpenseTracker.GetExpense(id)
			if err != nil {
				fmt.Printf("error in fetching expense: %s\n", err)
				return
			}
			exp.Print()
			return
		}
		if month > 12 || month < 0 {
			fmt.Println("invalid month enter")
			fmt.Println("valid values are [1, 12]")
			return
		}
		expenses, err := ExpenseTracker.ListExpense()
		if err != nil {
			fmt.Printf("error in fetching the expenses: %s\n", err)
			return
		}

		if month != 0 {
			var monthlyExpenses model.Expenses
			for _, exp := range expenses {
				if exp.CreatedAt.Month() == time.Month(month) {
					monthlyExpenses = append(monthlyExpenses, exp)
				}
			}
			expenses = monthlyExpenses
		}

		if year != 0 {
			var yearlyExpenses model.Expenses
			for _, exp := range expenses {
				if exp.CreatedAt.Year() == year {
					yearlyExpenses = append(yearlyExpenses, exp)
				}
			}
			expenses = yearlyExpenses
		}

		if cat != "" {
			cats := strings.Split(cat, ",")
			var categoricalExpenses model.Expenses
			for _, exp := range expenses {
				for _, category := range cats {
					if exp.Category == model.StringToCatEnum(strings.TrimSpace(category)) {
						categoricalExpenses = append(categoricalExpenses, exp)
					}
				}
			}
			expenses = categoricalExpenses
		}

		expenses.Print()
	},
}

func init() {
	listCmd.Flags().Int("id", 0, "expense id")
	listCmd.Flags().IntP("month", "m", 0, "month")
	listCmd.Flags().IntP("year", "y", 0, "year, by default it's current year")
	listCmd.Flags().StringP("category", "c", "", `expense categories, multiple categories can be passed comma separted
supported categries are: groceries, transport, medical, rent, entertainment`)

	rootCmd.AddCommand(listCmd)
}
