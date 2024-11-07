package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/DragoHex/expense-tracker/pkg/db"
	"github.com/DragoHex/expense-tracker/pkg/model"
	"github.com/DragoHex/expense-tracker/pkg/tracker"
)

var (
	ExpenseTracker *tracker.ExpenseTrackerImpl
	BudgetTracker  *tracker.BudgetRepoImpl

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "extr",
		Short: "A CLI tool for personal expense tracking",
		Long:  `A CLI tool for personal expense tracking`,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	dbFile := filepath.Join("artifacts", "db", "data.db")

	// Create the DB if it doesn't exists
	if _, err := os.Stat(dbFile); err != nil {
		log.Println("db file doesn't exist")
		log.Println("creating the db")
		if err := os.MkdirAll(filepath.Dir(dbFile), os.ModePerm); err != nil {
			log.Fatalf("Error creating directory: %v", err)
		}
	}

	// opne/initialise gorm DB
	dbLite, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("error in opening the db file: %s", err)
		return
	}

	// To create Expense DB table if it doensn't exist
	err = dbLite.AutoMigrate(&model.Expense{})
	if err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	}

	// To create Budget DB table if it doensn't exist
	err = dbLite.AutoMigrate(&model.Budget{})
	if err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	}

	dbRepo := db.NewDBRepoImpl(dbLite)
	dbBudgetRepo := db.NewDBRepoBudgetImpl(dbLite)

	ExpenseTracker = tracker.NewExpenseTrackerImpl(dbRepo)
	BudgetTracker = tracker.NewBudgetRepoImpl(dbBudgetRepo)
}
