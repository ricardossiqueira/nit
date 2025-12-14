/*Package db
 */
package db

import (
	"context"
	"database/sql"
	"fmt"

	"nit/internal/config"
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

func (s *Store) SaveRun(ctx context.Context, c *config.Config, r *llm.Response) error {
	query := `
		insert into runs (model, endpoint, prompt, response, duration_ms)
		values ('%s', '%s', '%s', '%s', %d);
	`
	finalQuery := fmt.Sprintf(query, c.Model.ModelName, c.Model.Endpoint, c.Prompt.SystemInstructions, r.Message.Content, r.TotalDuration)

	_, err := s.DB.Exec(finalQuery)
	if err != nil {
		return fmt.Errorf("error saving run: %w", err)
	}

	return nil
}
