package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// deleteBudgetCmd represents the budget command
var deleteBudgetCmd = &cobra.Command{
	Use:   "delete",
	Short: "A command to delete a budget entry.",
	Long: `A command to delete a budget entry.

If user doesn't provide any month or year.
Then the budget is deleted for the current month.
If a month is provided then it will delete for the current year.
User cannot use only year flag.
It has to be accompanies with month flag.'`,
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

		if cmd.Flags().Changed("year") && !cmd.Flags().Changed("month") {
			fmt.Println("only year passed, month is also needed")
			return
		}

		err = BudgetTracker.DeleteBudget(month, year)
		if err != nil {
			fmt.Printf("error in deleting budget entry: %s", err)
			return
		}
	},
}

func init() {
	t := time.Now()
	deleteBudgetCmd.Flags().IntP("month", "m", int(t.Month()), "month")
	deleteBudgetCmd.Flags().IntP("year", "y", t.Year(), "year")

	budgetCmd.AddCommand(deleteBudgetCmd)
}
