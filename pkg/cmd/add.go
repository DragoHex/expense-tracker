package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A command to add new expense",
	Long:  `A command to add new expense`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
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
