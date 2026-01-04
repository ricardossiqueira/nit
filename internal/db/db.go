/*Package db
 */
package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const createRunsTableSQL = `
CREATE TABLE IF NOT EXISTS runs
(
	id INTEGER PRIMARY KEY autoincrement,
	model TEXT,
	current_branch TEXT,
	endpoint TEXT,
	system_prompt TEXT,
	user_prompt TEXT,
	type TEXT,
	response TEXT,
	status_code INTEGER,
	duration_ms INTEGER,
	created_at DATETIME default (datetime('now'))
);
`

// Open opens a SQLite database at the provided path and verifies connectivity.
// Schema initialization is handled separately by InitSchema.
func Open(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("error initializing sqlite: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to sqlite: %w", err)
	}

	return db, nil
}

// InitSchema creates required tables if they do not exist.
func InitSchema(db *sql.DB) error {
	if db == nil {
		return fmt.Errorf("database connection is nil")
	}

	if _, err := db.Exec(createRunsTableSQL); err != nil {
		_ = db.Close()
		return fmt.Errorf("error running sql statement: %w", err)
	}

	return nil
}
