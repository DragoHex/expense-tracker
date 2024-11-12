package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// budgetCmd represents the budget command
var budgetCmd = &cobra.Command{
	Use:   "budget",
	Short: "A command to set monthly budget for the expenses.",
	Long: `A command to set monthly budget for the expenses.

If no sub-command or flag given it will print the budget for the current month.`,
	Run: func(cmd *cobra.Command, args []string) {
		t := time.Now()
		budg, err := BudgetTracker.GetBudget(int(t.Month()), t.Year())
		if err != nil {
			fmt.Printf("error in fetching the budget: %s\n", err)
			fmt.Println("could be because of no budget set for the month")
			fmt.Println("use budget -a <budget_amount> to set budget for the month")
			return
		}
		budg.Print()
	},
}

func init() {
	rootCmd.AddCommand(budgetCmd)
}
