# Go Starter

A minimalistic [GitHub repository template](https://docs.github.com/en/repositories/creating-and-managing-repositories/creating-a-repository-from-a-template) to kickstart your [Go](https://go.dev/) project

## Setup

Install [Go](https://go.dev/), [Lefthook](https://lefthook.dev/), [dprint](https://dprint.dev/), and [golangci-lint](https://golangci-lint.run/welcome/install/), then register the pre-commit hook:

```sh
lefthook install
```

## Development

Run the tests with:

```sh
go test ./... -cover
```

CI enforces 100% code coverage, with the thresholds set in `.testcoverage.yaml`.

Each `git commit` runs the pre-commit hook registered during setup, which checks your changes and fixes what it can in place. If it fails, fix any reported issues, re-stage the changed files, and commit again.

After committing, push to `main` or open a pull request from another branch — CI will run the same checks across all files, plus additional checks of its own.
