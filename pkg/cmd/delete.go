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
		fmt.Println("delete called")
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
