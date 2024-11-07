package cmd

import (
	"fmt"
	"time"

	"github.com/DragoHex/expense-tracker/pkg/model"
	"github.com/DragoHex/expense-tracker/pkg/utils"
	"github.com/spf13/cobra"
)

// listBudgetCmd represents the budget command
var listBudgetCmd = &cobra.Command{
	Use:   "list",
	Short: "A command to list monthly budget for the expenses.",
	Long: `A command to list monthly budget for the expenses.

If user doesn't provide any month or year.
Then the budget listed for the current month.
If a month is provided then it will for the current year.
If year is provided, then budget is fetched for all months of the year.'`,
	Run: func(cmd *cobra.Command, args []string) {
		month, err := cmd.Flags().GetInt("month")
		if err != nil {
			fmt.Printf("error in getting the flag value: %s\n", err)
			return
		}
		year, err := cmd.Flags().GetInt("year")
		if err != nil {
			fmt.Printf("error in getting the flag value: %s\n", err)
			return
		}

		if month < 1 || month > 12 {
			fmt.Printf("enter a valid month [1, 12]")
			return
		}

		if !cmd.Flags().Changed("month") && !cmd.Flags().Changed("year") {
			fmt.Println("Listing all the budget entries")
			budgets, err := BudgetTracker.ListBudget()
			if err != nil {
				fmt.Printf("error in fetching the budget entries: %s\n", err)
				return
			}
			budgets.Print()
			return
		}

		// If only year is passed
		if cmd.Flags().Changed("year") && !cmd.Flags().Changed("month") {
			fmt.Printf("Listing budget entries from year: %d\n", year)
			var filteredBudgets model.Budgets
			budgets, err := BudgetTracker.ListBudget()
			if err != nil {
				fmt.Printf("error in fetching the budget entries: %s\n", err)
				return
			}
			for _, budg := range budgets {
				_, y, err := utils.SplitStringToTime(budg.MonthYear)
				if err != nil {
					fmt.Printf("error in splitting month year: %s", err)
					return
				}
				if y == year {
					filteredBudgets = append(filteredBudgets, budg)
				}
			}
			filteredBudgets.Print()
			return
		}

		budg, err := BudgetTracker.GetBudget(month, year)
		if err != nil {
			fmt.Printf("error in fetching the budget entries: %s\n", err)
			return
		}
		budg.Print()
	},
}

func init() {
	t := time.Now()
	listBudgetCmd.Flags().IntP("month", "m", int(t.Month()), "month")
	listBudgetCmd.Flags().IntP("year", "y", t.Year(), "year")

	budgetCmd.AddCommand(listBudgetCmd)
}
