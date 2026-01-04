/*Package db
 */
package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) (*Store, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	return &Store{db: db}, nil
}

type Run struct {
	Model         string
	CurrentBranch string
	Endpoint      string
	SystemPrompt  string
	UserPrompt    string
	Type          string
	Response      string
	DurationMS    int64
}

type Draft struct {
	PRTitle       string
	PRDescription string
	CommitMessage string
}

const (
	insertRunQuery = `
INSERT INTO runs (model, current_branch, endpoint, system_prompt, user_prompt, type, response, duration_ms)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);`

	draftByIDQuery = `
SELECT 
	json_extract(response, '$.pr_title') as pr_title,
	json_extract(response, '$.pr_description') as pr_description,
	json_extract(response, '$.commit_message') as commit_message
FROM runs 
WHERE id = ?
LIMIT 1;`

	lastDraftQuery = `
SELECT 
	json_extract(response, '$.pr_title') as pr_title,
	json_extract(response, '$.pr_description') as pr_description,
	json_extract(response, '$.commit_message') as commit_message
FROM runs
ORDER BY id DESC
LIMIT 1;`

	lastByTypeQuery = `
SELECT response
FROM runs
WHERE type = ?
ORDER BY id DESC
LIMIT 1;`
)

func (s *Store) SaveRun(ctx context.Context, r *Run) error {
	if s == nil || s.db == nil {
		return fmt.Errorf("store is not initialized")
	}

	_, err := s.db.ExecContext(
		ctx,
		insertRunQuery,
		r.Model,
		r.CurrentBranch,
		r.Endpoint,
		r.SystemPrompt,
		r.UserPrompt,
		r.Type,
		r.Response,
		r.DurationMS,
	)
	if err != nil {
		return fmt.Errorf("error saving run: %w", err)
	}

	return nil
}

func (s *Store) GetDraftByID(ctx context.Context, id string) (*Draft, error) {
	if s == nil || s.db == nil {
		return nil, fmt.Errorf("store is not initialized")
	}

	row := s.db.QueryRowContext(ctx, draftByIDQuery, id)

	draft, err := scanDraft(row)
	if err != nil {
		return nil, fmt.Errorf("draft %s not found: %w", id, err)
	}

	return draft, nil
}

func (s *Store) GetLastDraft(ctx context.Context) (*Draft, error) {
	if s == nil || s.db == nil {
		return nil, fmt.Errorf("store is not initialized")
	}

	row := s.db.QueryRowContext(ctx, lastDraftQuery)

	draft, err := scanDraft(row)
	if err != nil {
		return nil, fmt.Errorf("draft not found: %w", err)
	}

	return draft, nil
}

func (s *Store) GetLastByType(ctx context.Context, runType string) (string, error) {
	if s == nil || s.db == nil {
		return "", fmt.Errorf("store is not initialized")
	}

	row := s.db.QueryRowContext(ctx, lastByTypeQuery, runType)

	var resp string
	if err := row.Scan(&resp); err != nil {
		return "", fmt.Errorf("run of type %s not found: %w", runType, err)
	}

	return resp, nil
}

func scanDraft(row *sql.Row) (*Draft, error) {
	var draft Draft
	if err := row.Scan(&draft.PRTitle, &draft.PRDescription, &draft.CommitMessage); err != nil {
		return nil, err
	}
	return &draft, nil
}
