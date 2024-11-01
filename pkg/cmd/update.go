package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A command to update existing expenses",
	Long:  `A command to update existing expenses`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
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
