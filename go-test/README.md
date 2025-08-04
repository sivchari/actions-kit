# Go Test Action

Runs Go tests with race detection and shuffle enabled.

## Usage

```yaml
- name: Run Go tests
  uses: sivchari/actions-kit/go-test@main
  with:
    go-version-file: 'go.mod'  # optional, default: 'go.mod'
    working-directory: '.'     # optional, default: '.'
    test-args: '-coverprofile=coverage.out'  # optional
```

## Inputs

| Input | Description | Required | Default |
|-------|-------------|----------|---------|
| `go-version-file` | Path to go.mod file for Go version detection | No | `go.mod` |
| `working-directory` | Working directory to run tests in | No | `.` |
| `test-args` | Additional arguments to pass to go test | No | `''` |

## Example

```yaml
name: Test
on:
  pull_request:

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Run tests
        uses: sivchari/actions-kit/go-test@main
        with:
          test-args: '-coverprofile=coverage.out -timeout=10m'
```