package cmd

import (
	"database/sql"
	"log"
)

// RunQuery processes db.Query into readily usable cols and rows
func RunQuery(db *sql.DB, query string, args ...any) (cols []string, rows [][]any) {
	qRows, err := db.Query(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer qRows.Close()

	cols, err = qRows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	for qRows.Next() {
		ptrs := make([]any, len(cols))
		row := make([]any, len(cols))
		for i := range ptrs {
			ptrs[i] = &row[i]
		}
		if err := qRows.Scan(ptrs...); err != nil {
			log.Fatal(err)
		}
		rows = append(rows, row)
	}
	return
}
