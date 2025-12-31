/*Package db
 */
package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Open(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("error initializing sqlite: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to sqlite: %w", err)
	}

	query := `
		CREATE TABLE IF NOT EXISTS runs
			(
				id INTEGER PRIMARY KEY autoincrement,
				model TEXT,
				current_branch TEXT,
				endpoint TEXT,
				prompt TEXT,
				response	JSON,
				status_code INTEGER,
				duration_ms INTEGER,
				created_at DATETIME default (datetime('now'))
			);
	`

	if _, err := db.Exec(query); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("error running sql statement: %w", err)
	}

	return db, nil
}
