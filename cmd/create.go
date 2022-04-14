package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a sqlite3 db file",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		mkdirs, _ := cmd.Flags().GetBool("mkdirs")
		if mkdirs {
			os.MkdirAll(filepath.Dir(dbFile), 0755)
		}

		if _, err := os.Create(dbFile); err != nil {
			log.Fatal(err)
		}
		fmt.Println("empty db file created at", dbFile)
	},
}

func init() {
	createCmd.Flags().Bool("mkdirs", false, "create missing directories")

	rootCmd.AddCommand(createCmd)
}
