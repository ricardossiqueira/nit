# nit

`nit` is a lightweight Go CLI that uses **local LLMs via Ollama** to generate **Git commit messages** and **pull request content** directly from your local Git diff.

It is local-first, scriptable, and designed to fit naturally into an existing Git workflow.

---

## Requirements

* Go 1.21+
* Git
* **Ollama installed and available in PATH**

---

## Installation

```sh
git clone https://github.com/ricardossiqueira/nit.git
cd nit
go build
```

Or run directly:

```sh
go run .
```

---

## Initial Configuration (Required)

Before using `nit`, generate the initial config file.

```sh
go run . config init
```

This creates a configuration file defining the LLM provider and model.

### Check configured model

```sh
go run . config show -m
```

### Install the configured model with Ollama

```sh
ollama pull $(go run . config show -m)
```

Make sure Ollama is running:

```sh
ollama serve
```

---

## Key Features

* Generate draft **commit messages** from `git diff`
* Generate **pull request title and body** from `git diff`
* Structured, deterministic Markdown prompts
* Local persistence using SQLite
* Fully local execution (no cloud APIs)
* Designed for aliases and automation

---

## Usage

### Commit message draft

```sh
go run . draft commit
```

---

### Pull request title and body

```sh
go run . draft pr
```

This command generates both the **PR title** and **PR body**.

---

## Recommended Daily Workflow

`nit` is most effective when composed with existing Git commands.

### Commit alias

```sh
alias ncm="git commit -m '$(go run . draft -l commit)'"
```

Usage:

```sh
ncm
```

---

### Pull request alias (GitHub CLI)

```sh
alias npr="gh pr create \
  --title '$(go run . draft -l pr --title)' \
  --body '$(go run . draft -l pr --body)'"
```

> `draft pr` is executed once per output type, ensuring the title and body are generated consistently from the same diff.

Usage:

```sh
npr
```

---

## Design Principles

* Local-first: Git is the source of truth
* Drafts, not final decisions
* Explicit prompts with predictable output
* No hidden automation or side effects

---

## License

MIT
