package cmd

import (
	"fmt"

	"github.com/DragoHex/expense-tracker/pkg/model"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A command to update existing expenses",
	Long:  `A command to update existing expenses`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			fmt.Printf("error in getting the flag value: %s\n", err)
			return
		}
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

		exp, err := ExpenseTracker.GetExpense(id)
		if err != nil {
			fmt.Printf("error in fetching expense with ID - %d: %s\n", id, err)
			return
		}

		if des == "" {
			des = exp.Description
		}
		if cat == "" {
			cat = model.Category(exp.Category).String()
		}
		if amount == 0 {
			amount = exp.Amount
		}

		err = ExpenseTracker.UpdateExpense(id, des, amount, model.StringToCatEnum(cat))
		if err != nil {
			fmt.Printf("error updating expense with ID - %d: %s\n", id, err)
			return
		}

		fmt.Println("Updated expense from:")
		exp.Print()
		fmt.Println("\nTo:")
		model.Expense{
			ID:          id,
			Description: des,
			Amount:      amount,
			Category:    model.StringToCatEnum(cat),
		}.Print()
	},
}

func init() {
	updateCmd.Flags().Int("id", 0, "expense id")
	updateCmd.Flags().StringP("description", "d", "", "expense description")
	updateCmd.Flags().StringP("category", "c", "", "expense category")
	updateCmd.Flags().IntP("amount", "a", 0, "expense amount")

	err := updateCmd.MarkFlagRequired("id")
	if err != nil {
		fmt.Printf("error in setting flag as required: %s", err)
		return
	}

	rootCmd.AddCommand(updateCmd)
}
