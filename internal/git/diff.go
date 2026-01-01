/*
Package git
*/
package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type DiffContext struct {
	BaseBranch   string
	FilesChanged []string
	Summary      string
	RawDiff      string
}

func ParseDiff(baseBranch string, maxLines int) (*DiffContext, error) {
	var cmdOut bytes.Buffer
	var cmdErr bytes.Buffer

	// files changed
	cmd := exec.Command("git", "diff", "--name-only", baseBranch)
	cmd.Stdout = &cmdOut
	cmd.Stderr = &cmdErr
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("git diff --name-only: %s", cmdErr.String())
	}

	var files []string
	for line := range strings.SplitSeq(cmdOut.String(), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		//TODO: make lang dynamic
		// if strings.HasSuffix(line, ".go") {
		// 	files = append(files, line)
		// }
	}

	//TODO: support diff files
	// diff
	diffCmd := exec.Command("git", "diff", baseBranch, "--", "*")
	diffOut, err := diffCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("git diff: %w", err)
	}

	diffLines := strings.Split(string(diffOut), "\n")
	if maxLines > 0 && len(diffLines) > maxLines {
		diffLines = diffLines[:maxLines]
	}

	rawDiff := strings.Join(diffLines, "\n")

	// simple summary
	summary := buildSummary(files)

	return &DiffContext{
		BaseBranch:   baseBranch,
		FilesChanged: files,
		Summary:      summary,
		RawDiff:      rawDiff,
	}, nil
}

func buildSummary(files []string) string {
	if len(files) == 0 {
		return "No changes detected."
	}

	var buf bytes.Buffer
	buf.WriteString("Files changed:\n")
	for _, f := range files {
		buf.WriteString("- " + f + "\n")
	}
	return buf.String()
}
