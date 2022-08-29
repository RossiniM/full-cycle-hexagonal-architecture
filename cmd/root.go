/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"database/sql"
	"github.com/RossiniM/full-cycle-hexagonal-architecture/adapters/db"
	"github.com/RossiniM/full-cycle-hexagonal-architecture/application"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var db1, _ = sql.Open("sqlite3", "productDB.db")

func createTable(db2 *sql.DB) {
	table := ` create table products(id string, name string, price float,status string);`
	statement, err := db2.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}

var productDB = db.New(db1)
var productService = application.ProductService{Persistence: productDB}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "full-cycle-hexagonal-architecture",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.full-cycle-hexagonal-architecture.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
