package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A command to list expenses",
	Long:  `A command to list expenses`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	listCmd.Flags().Int("id", 0, "expense id")
	listCmd.Flags().IntP("month", "m", 1, "month")
	listCmd.Flags().StringP("category", "c", "", `expense categories, multiple categories can be passed comma separted
supported categries are: groceries, transport, medical, rent, entertainment`)

	rootCmd.AddCommand(listCmd)
}
