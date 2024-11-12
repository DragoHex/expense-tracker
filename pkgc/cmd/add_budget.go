package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// addBudgetCmd represents the budget command
var addBudgetCmd = &cobra.Command{
	Use:   "add",
	Short: "A command to add monthly budget for the expenses.",
	Long: `A command to add monthly budget for the expenses.

If user doesn't provide any month or year.
Then the budget is set for the current month.
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
			fmt.Println("enter a valid month [1, 12]")
			return
		}

		if amount < 0 {
			fmt.Println("budget cannot be negative")
		}

		if cmd.Flags().Changed("year") && !cmd.Flags().Changed("month") {
			fmt.Println("only year passed, month is also needed")
			return
		}

		budg, err := BudgetTracker.CreateBudget(month, year, amount)
		if err != nil {
			fmt.Printf("error in setting budget: %s\n", err)
			return
		}
		fmt.Printf("Budget added successfully for %s-%d\n\n", time.Month(month).String(), year)
		budg.Print()
	},
}

func init() {
	t := time.Now()
	addBudgetCmd.Flags().IntP("month", "m", int(t.Month()), "month")
	addBudgetCmd.Flags().IntP("year", "y", t.Year(), "year")
	addBudgetCmd.Flags().IntP("amount", "a", 0, "budget amount")

	err := addBudgetCmd.MarkFlagRequired("amount")
	if err != nil {
		fmt.Printf("error in setting flag as required: %s", err)
		return
	}

	budgetCmd.AddCommand(addBudgetCmd)
}
