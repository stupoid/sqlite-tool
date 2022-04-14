package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jedib0t/go-pretty/table"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "ls quickly list data from the db",
	Long: `If no options are passed, all tables in the db will be listed.
If a table is passed as an option, the top 20 rows of that table will be listed instead`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		query := "SELECT name FROM sqlite_master WHERE type='table'"

		tableName, _ := cmd.Flags().GetString("table")
		if tableName != "" {
			query = fmt.Sprintf("SELECT * FROM %s LIMIT 20", tableName)
		}

		db, err := sql.Open("sqlite3", dbFile)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		cols, rows := RunQuery(db, query)

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)

		tableRow := make([]interface{}, len(cols))
		for i, col := range cols {
			tableRow[i] = col
		}
		t.AppendHeader(tableRow)
		for _, row := range rows {
			t.AppendRow(row)
		}
		fmt.Println(query)
		t.Render()
	},
}

func init() {
	lsCmd.Flags().StringP("table", "t", "", "list top 20 rows from table")

	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
