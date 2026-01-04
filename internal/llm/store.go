/* Package llm
 */
package llm

import (
	"context"
	"time"
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
	SaveRun(ctx context.Context, run *Run) error
}
