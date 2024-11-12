package cmd

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/DragoHex/expense-tracker/pkgc/db"
	"github.com/DragoHex/expense-tracker/pkgc/tracker"
)

var (
	dbQueries      *db.Queries
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

	// initialise gorm DB
	dbObj, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("error in opening the db file: %s", err)
		return
	}

	dbLite, err := dbObj.DB()
	if err != nil {
		log.Fatalf("error in getting *sql.DB from *gorm.DB: %s", err)
		return
	}
	err = initTables(dbLite)
	if err != nil {
		log.Fatalf("error in intialising the tables: %s", err)
		return
	}

	ExpenseTracker = tracker.NewExpenseTrackerImpl(dbLite)
	BudgetTracker = tracker.NewBudgetRepoImpl(dbLite)
}

func initTables(db *sql.DB) error {
	// create budget table
	createTableQuery := `
CREATE TABLE IF NOT EXISTS budget (
    month_year PRIMARY KEY,
    amount INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`
	_, err := db.Exec(createTableQuery)
	if err != nil {
		return err
	}

	// create expense table
	createTableQuery = `
CREATE TABLE IF NOT EXISTS expense (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	description TEXT,
	amount INTEGER,
	category INTEGER,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`
	_, err = db.Exec(createTableQuery)
	return err
}
