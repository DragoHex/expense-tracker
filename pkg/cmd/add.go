package cmd

import (
	"fmt"
	"time"

	"github.com/DragoHex/expense-tracker/pkg/model"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A command to add new expense",
	Long:  `A command to add new expense`,
	Run: func(cmd *cobra.Command, args []string) {
		des, err := cmd.Flags().GetString("description")
		if err != nil {
			fmt.Printf("error in getting the flag value: %s\n", err)
			return
		}
		cat, err := cmd.Flags().GetString("category")
		if err != nil {
			fmt.Printf("error in getting the flag value: %s\n", err)
			return
		}
		amount, err := cmd.Flags().GetInt("amount")
		if err != nil {
			fmt.Printf("error in getting the flag value: %s\n", err)
			return
		}

		if amount < 0 {
			fmt.Println("expense cannot be negative")
			return
		}

		exp, err := ExpenseTracker.CreateExpense(des, amount, model.StringToCatEnum(cat))
		if err != nil {
			fmt.Printf("error in adding expense: %s\n", err)
			return
		}

		t := time.Now()
		month := int(t.Month())
		year := t.Year()

		exps, _ := ExpenseTracker.ListFilteredExpense(month, year)
		bud, _ := BudgetTracker.GetBudget(month, year)

		if exps.Total() > bud.Amount {
			fmt.Println()
			fmt.Printf("\033[0;31mEpenses has crossed this month's budget: %d by %d \033[0m\n", bud.Amount, exps.Total()-bud.Amount)
			fmt.Println()
		}

		fmt.Printf("Expense added successfully (ID:%d)\n\n", exp.ID)
		exp.Print()
	},
}

func init() {
	addCmd.Flags().StringP("description", "d", "", "expense description")
	addCmd.Flags().StringP("category", "c", "", "expense category, valid values: groceries, transport, medical, rent, entertainment")
	addCmd.Flags().IntP("amount", "a", 0, "expense amount")

	err := addCmd.MarkFlagRequired("description")
	if err != nil {
		fmt.Printf("error in setting flag as required: %s", err)
		return
	}
	err = addCmd.MarkFlagRequired("amount")
	if err != nil {
		fmt.Printf("error in setting flag as required: %s", err)
		return
	}

	rootCmd.AddCommand(addCmd)
}
