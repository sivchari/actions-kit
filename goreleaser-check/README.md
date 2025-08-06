# GoReleaser Check Action

Runs GoReleaser check and snapshot build to validate configuration.

## Usage

```yaml
- name: Run GoReleaser check
  uses: sivchari/actions-kit/goreleaser-check@main
  with:
    go-version-file: 'go.mod'  # optional, default: 'go.mod'
    working-directory: '.'     # optional, default: '.'
    goreleaser-version: '~> v2'  # optional, default: '~> v2'
    args: ''                   # optional
    snapshot-build: 'true'     # optional, default: 'true'
```

## Inputs

| Input | Description | Required | Default |
|-------|-------------|----------|---------|
| `go-version-file` | Path to go.mod file for Go version detection | No | `go.mod` |
| `working-directory` | Working directory to run GoReleaser check in | No | `.` |
| `goreleaser-version` | Version of GoReleaser to use | No | `~> v2` |
| `args` | Additional arguments to pass to GoReleaser check | No | `''` |
| `snapshot-build` | Run snapshot build after check | No | `true` |

## Example

```yaml
name: Release Check
on:
  pull_request:

jobs:
  release-check:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0
      - name: Run GoReleaser check
        uses: sivchari/actions-kit/goreleaser-check@main
        with:
          args: '--verbose'
```