package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	dbFile  string
	dsn     string // ref: https://github.com/mattn/go-sqlite3#connection-string
	rootCmd = &cobra.Command{
		Use:   "sqlite-tool",
		Short: "sqlite-tool is a way to manage your SQLite DBs on the CLI",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&dbFile, "db", "", "sqlite3 db file to access (required)")
	rootCmd.MarkFlagRequired("db")
}

func initConfig() {
	if dbFile == "" {
		fmt.Println("db file not provided")
		os.Exit(1)
	}
}
