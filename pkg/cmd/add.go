package cmd

import (
	"fmt"

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

		exp, err := ExpenseTracker.CreateExpense(des, amount, model.StringToCatEnum(cat))
		if err != nil {
			fmt.Printf("error in adding expense: %s\n", err)
			return
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
