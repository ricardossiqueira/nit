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

func (s *Store) GetDraftByID(ctx context.Context, id string) (*llm.DraftOutput, error) {
	query := `
        SELECT 
            json_extract(response, '$.pr_title') as pr_title,
            json_extract(response, '$.pr_description') as pr_description,
            json_extract(response, '$.commit_message') as commit_message
        FROM runs 
        WHERE id = ?
        LIMIT 1
    `

	row := s.DB.QueryRowContext(ctx, query, id)

	var draft llm.DraftOutput
	err := row.Scan(&draft.PRTitle, &draft.PRDescription, &draft.CommitMessage)
	if err != nil {
		return nil, fmt.Errorf("draft %s not found: %w", id, err)
	}

	return &draft, nil
}

func (s *Store) GetLastDraft(ctx context.Context) (*llm.DraftOutput, error) {
	query := `
        SELECT 
            json_extract(response, '$.pr_title') as pr_title,
            json_extract(response, '$.pr_description') as pr_description,
            json_extract(response, '$.commit_message') as commit_message
        FROM runs
        ORDER BY id DESC
        LIMIT 1
    `

	row := s.DB.QueryRowContext(ctx, query)

	var draft llm.DraftOutput
	err := row.Scan(&draft.PRTitle, &draft.PRDescription, &draft.CommitMessage)
	if err != nil {
		return nil, fmt.Errorf("draft not found: %w", err)
	}

	return &draft, nil
}
