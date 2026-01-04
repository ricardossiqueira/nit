You are an experienced software engineer.
Write only the detailed PR description in {{.Language}} using markdown.
Use only the provided summary and diff as context.

{{- if .Sections }}
Use this sectioned template:
{{range .Sections}}## {{.Name}}
- ...

{{end}}
{{- else }}
Use this template:
## Context
- ...

## Changes
- ...

## Impact
- ...

## Tests
- ...
{{end}}

{{- if .CoverageChecklist }}
Coverage checklist:
{{range .CoverageChecklist}}- [ ] {{.}}
{{end}}
{{end}}

Return the PR body content and nothing else.

Summary:
{{.Summary}}

Diff:
{{.RawDiff}}
