/*Package db
 */
package db

import (
	"context"
	"database/sql"
	"fmt"

	"nit/internal/llm"
)

type Store struct {
	DB *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		DB: db,
	}
}

func (s *Store) SaveRun(ctx context.Context, r *llm.Run) error {
	query := `
  	INSERT INTO runs (model, current_branch, endpoint, prompt, response, duration_ms)
  	VALUES (?, ?, ?, ?, ?, ?);
	`

	responseJSON, err := r.Response.Marshal()
	if err != nil {
		return err
	}

	_, err = s.DB.ExecContext(
		ctx,
		query,
		r.Model,
		r.CurrentBranch,
		r.Endpoint,
		r.Prompt,
		responseJSON,
		r.DurationMS,
	)
	if err != nil {
		return fmt.Errorf("error saving run: %w", err)
	}

	return nil
}
