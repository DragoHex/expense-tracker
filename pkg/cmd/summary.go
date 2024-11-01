package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "A command to get expense summary",
	Long:  `A command to get expense summary`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("summary called")
	},
}

func init() {
	summaryCmd.Flags().IntP("month", "m", 1, "month")
	rootCmd.AddCommand(summaryCmd)
}
