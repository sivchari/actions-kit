# Go Mod Check Action

Checks if go.mod and go.sum are tidy and up-to-date using `go mod tidy -diff`.

## Usage

```yaml
- name: Check go modules
  uses: sivchari/actions-kit/go-mod-check@main
  with:
    go-version-file: 'go.mod'    # optional, default: 'go.mod'
    working-directory: '.'       # optional, default: '.'
    check-tools: 'true'          # optional, default: 'true'
```

## Inputs

| Input | Description | Required | Default |
|-------|-------------|----------|---------|
| `go-version-file` | Path to go.mod file for Go version detection | No | `go.mod` |
| `working-directory` | Working directory to run mod check in | No | `.` |
| `check-tools` | Also check tools/go.mod if it exists | No | `true` |

## What it checks

1. **Main module**: Runs `go mod tidy -diff` in the working directory
2. **Tools module** (if enabled): Runs `go mod tidy -diff` in `tools/` directory if it exists

The action will fail if any modules are not tidy (i.e., if `go mod tidy` would make changes).

## Example

```yaml
name: Go Mod Check
on:
  pull_request:

jobs:
  go-mod-check:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Check go modules
        uses: sivchari/actions-kit/go-mod-check@main
        with:
          check-tools: 'false'  # Skip tools check
```