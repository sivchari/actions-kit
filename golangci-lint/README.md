# golangci-lint Action

Runs golangci-lint with configurable options.

## Usage

```yaml
- name: Run golangci-lint
  uses: sivchari/actions-kit/golangci-lint@main
  with:
    go-version-file: 'go.mod'        # optional, default: 'go.mod'
    working-directory: '.'           # optional, default: '.'
    golangci-lint-version: 'latest'  # optional, default: 'latest'
    args: '--enable-all'             # optional
    timeout: '5m'                    # optional, default: '5m'
```

## Inputs

| Input | Description | Required | Default |
|-------|-------------|----------|---------|
| `go-version-file` | Path to go.mod file for Go version detection | No | `go.mod` |
| `working-directory` | Working directory to run linter in | No | `.` |
| `golangci-lint-version` | Version of golangci-lint to use | No | `latest` |
| `args` | Additional arguments to pass to golangci-lint | No | `''` |
| `timeout` | Timeout for golangci-lint execution | No | `5m` |

## Example

```yaml
name: Lint
on:
  pull_request:

jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Run golangci-lint
        uses: sivchari/actions-kit/golangci-lint@main
        with:
          golangci-lint-version: 'v1.55.2'
          args: '--enable-all --disable=gochecknoglobals'
          timeout: '10m'
```