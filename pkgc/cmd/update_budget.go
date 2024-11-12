package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// updateBudgetCmd represents the budget command
var updateBudgetCmd = &cobra.Command{
	Use:   "update",
	Short: "A command to update monthly budget for the expenses.",
	Long: `A command to update monthly budget for the expenses.

If user doesn't provide any month or year.
Then the budget is update for the current month.
If a month is provided then it will for the current year.
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
		amount, err := cmd.Flags().GetInt("amount")
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

		err = BudgetTracker.UpdateBudget(month, year, amount)
		if err != nil {
			fmt.Printf("error in updating the budget entry: %s", err)
			return
		}
		fmt.Printf("budget updated for the month %s to %d\n", time.Month(month).String(), amount)
	},
}

func init() {
	t := time.Now()
	updateBudgetCmd.Flags().IntP("month", "m", int(t.Month()), "month")
	updateBudgetCmd.Flags().IntP("year", "y", t.Year(), "year")
	updateBudgetCmd.Flags().IntP("amount", "a", 0, "budget amount")

	err := updateBudgetCmd.MarkFlagRequired("amount")
	if err != nil {
		fmt.Printf("error in setting flag as required: %s", err)
		return
	}
	budgetCmd.AddCommand(updateBudgetCmd)
}
