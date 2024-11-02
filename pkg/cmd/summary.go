package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/DragoHex/expense-tracker/pkg/model"
	"github.com/spf13/cobra"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "A command to get expense summary",
	Long:  `A command to get expense summary`,
	Run: func(cmd *cobra.Command, args []string) {
		month, err := cmd.Flags().GetInt("month")
		if err != nil {
			fmt.Printf("error in getting the flag: %s\n", err)
			return
		}
		year, err := cmd.Flags().GetInt("year")
		if err != nil {
			fmt.Printf("error in getting the flag: %s\n", err)
			return
		}

		if month > 12 || month < 0 {
			fmt.Println("invalid month enter")
			fmt.Println("valid values are [1, 12]")
			return
		}

		cat, err := cmd.Flags().GetString("category")
		if err != nil {
			fmt.Printf("error in getting the flag value: %s", err)
			return
		}

		exps, err := ExpenseTracker.ListExpense()
		if err != nil {
			fmt.Printf("error in fetching all the expenses: %s\n", err)
			return
		}

		if cat != "" {
			cats := strings.Split(cat, ",")
			var categoricalExpenses model.Expenses
			for _, exp := range exps {
				for _, category := range cats {
					if exp.Category == model.StringToCatEnum(strings.TrimSpace(category)) {
						categoricalExpenses = append(categoricalExpenses, exp)
					}
				}
			}
			exps = categoricalExpenses
		}
		exps.Summary(month, year)
	},
}

func init() {
	t := time.Now()
	summaryCmd.Flags().IntP("month", "m", 0, "month")
	summaryCmd.Flags().IntP("year", "y", t.Year(), "year, by default it's current year")
	summaryCmd.Flags().StringP("category", "c", "", `expense categories, multiple categories can be passed comma separted
supported categries are: groceries, transport, medical, rent, entertainment`)
	rootCmd.AddCommand(summaryCmd)
}
