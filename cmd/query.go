package cmd

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jedib0t/go-pretty/table"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use: "query",
	Short: "Query executes a query that returns rows, typically a SELECT. The args are for any placeholder parameters in the query.	",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		validOutputs := map[string]struct{}{
			"":     {},
			"csv":  {},
			"html": {},
			"md":   {},
			"json": {},
		}

		if _, ok := validOutputs[output]; !ok {
			fmt.Println("invalid ouput option")
			os.Exit(1)
		}

		db, err := sql.Open("sqlite3", dbFile)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		dbArgs := make([]any, len(args)-1)
		for i, arg := range args[1:] {
			dbArgs[i] = arg
		}

		cols, rows := RunQuery(db, args[0], dbArgs...)

		if output == "json" {
			mapArray := []interface{}{}
			for _, row := range rows {
				mappedData := map[string]interface{}{}
				for i := range cols {
					mappedData[cols[i]] = row[i]
				}
				mapArray = append(mapArray, mappedData)
			}
			jsonBytes, err := json.MarshalIndent(mapArray, "  ", "  ")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(jsonBytes))
			os.Exit(0)
		}

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
		switch output {
		case "csv":
			t.RenderCSV()
		case "html":
			t.RenderHTML()
		case "md":
			t.RenderMarkdown()
		default:
			t.Render()
		}
	},
}

func init() {
	queryCmd.Flags().StringP("output", "o", "", "output of the query results. (csv|html|md|json)")

	rootCmd.AddCommand(queryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
