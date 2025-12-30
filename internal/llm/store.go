/* Package llm
 */
package llm

import (
	"context"
	"time"
)

type Run struct {
	Model      string
	Endpoint   string
	System     string
	Prompt     string
	Response   string
	DurationMS int64
	CreatedAt  time.Time
}

type RunStore interface {
	SaveRun(ctx context.Context, run *Run) error
}
