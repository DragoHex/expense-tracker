package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/DragoHex/expense-tracker/pkgc/db"
	"github.com/DragoHex/expense-tracker/pkgc/utils"
	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "A command to export the expense data to CSV",
	Long:  `A command to export the expense data to CSV`,
	Run: func(cmd *cobra.Command, args []string) {
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Printf("error in getting the flag value: %s\n", err)
			return
		}

		err = utils.CreateFile(output)
		if err != nil {
			fmt.Printf("error in creating csv file: %s\n", err)
			return
		}

		err = exportCSV(output)
		if err != nil {
			fmt.Printf("error in exporting the expense data: %s\n", err)
			return
		}
		fmt.Printf("expense data exported to %s\n", output)
	},
}

func init() {
	exportCmd.Flags().StringP("output", "o", "expense.csv", "export file path")
	rootCmd.AddCommand(exportCmd)
}

// export expense data to csv file
func exportCSV(output string) error {
	file, err := os.OpenFile(output, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// add headers
	header := []string{"ID", "Created At", "Updated At", "Description", "Amount", "Category"}
	err = writer.Write(header)
	if err != nil {
		return err
	}

	expenses, err := ExpenseTracker.ListExpense()
	if err != nil {
		return err
	}

	// write data rows
	for _, exp := range expenses {
		row := []string{
			strconv.Itoa(exp.ID),
			exp.CreatedAt.Format("2006-01-02"),
			exp.UpdatedAt.Format("2006-01-02"),
			exp.Description,
			strconv.Itoa(exp.Amount),
			db.Category(exp.Category).String(),
		}

		err := writer.Write(row)
		if err != nil {
			return err
		}
	}

	return nil
}
