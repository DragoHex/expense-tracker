package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A command to delete existing expense",
	Long:  `A command to delete existing expense`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			fmt.Printf("error in getting the flag value: %s\n", err)
			return
		}

		exp, err := ExpenseTracker.GetExpense(id)
		if err != nil {
			fmt.Println("expense not found")
			return
		}

		err = ExpenseTracker.DeleteExpense(id)
		if err != nil {
			fmt.Printf("error in deleting expense with ID - %d: %s\n", id, err)
			return
		}
		fmt.Printf("Deleted expense with ID: %d\n", id)
		exp.Print()
	},
}

func init() {
	deleteCmd.Flags().Int("id", 0, "expense id")

	err := deleteCmd.MarkFlagRequired("id")
	if err != nil {
		return
	}

	rootCmd.AddCommand(deleteCmd)
}
