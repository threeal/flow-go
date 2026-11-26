# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## About This Repository

A Go package for chaining tasks asynchronously in a flow. Its only code so far is the placeholder Fibonacci implementation in `pkg/myfibonacci/` inherited from `threeal/go-starter`, with the module path still `github.com/threeal/go-starter`.

## Gotchas

- `lefthook run pre-commit` skips every job when nothing is staged, even ones that ignore the staged file list — pass `--all-files` to run it outside an actual commit.
- The pre-commit hook fixes files in place, and the run fails if any file changes. Report that failure and leave the fixes for the user to review and re-stage before committing again.
- `go test` doesn't check coverage — only CI applies the 100% threshold, through the go-test-coverage action with the rules in `.testcoverage.yaml`.
- `.gitignore` ignores every dotfile and dotdir (`.*`) — a new one needs an explicit `!` exception there or it silently won't be tracked.
