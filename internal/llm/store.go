/* Package llm
 */
package llm

import (
	"context"
	"time"

	"nit/internal/db"
)

type Run struct {
	Model         string
	Endpoint      string
	SystemPrompt  string
	UserPrompt    string
	Response      string
	DurationMS    int64
	CreatedAt     time.Time
	CurrentBranch string
	Type          string
}

type RunStore interface {
	SaveRun(ctx context.Context, run *db.Run) error
}

// ToDBRun converts an llm.Run into a db.Run for persistence.
func ToDBRun(r *Run) *db.Run {
	if r == nil {
		return nil
	}

	return &db.Run{
		Model:         r.Model,
		CurrentBranch: r.CurrentBranch,
		Endpoint:      r.Endpoint,
		SystemPrompt:  r.SystemPrompt,
		UserPrompt:    r.UserPrompt,
		Type:          r.Type,
		Response:      r.Response,
		DurationMS:    r.DurationMS,
	}
}
