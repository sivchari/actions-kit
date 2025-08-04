# actions-kit

Reusable GitHub Actions for Go projects.

## Available Actions

### go-test
Runs Go tests with race detection and shuffle enabled.

```yaml
- uses: sivchari/actions-kit/go-test@main
```

### golangci-lint
Runs golangci-lint with configurable options and optional format checking.

```yaml
- uses: sivchari/actions-kit/golangci-lint@main
```

### go-mod-check
Checks if go.mod and go.sum are tidy and up-to-date.

```yaml
- uses: sivchari/actions-kit/go-mod-check@main
```

## Usage Examples

### Complete CI Workflow (Recommended)

Use the root action to run both tests and linting:

```yaml
name: CI
on:
  pull_request:

jobs:
  ci:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: sivchari/actions-kit@main
```

### Individual Actions

You can also use individual actions:

```yaml
name: CI
on:
  pull_request:

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: sivchari/actions-kit/go-test@main

  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: sivchari/actions-kit/golangci-lint@main
```

### Configuration Options

The root action supports various configuration options:

```yaml
- uses: sivchari/actions-kit@main
  with:
    go-version-file: 'go.mod'           # Path to go.mod (default)
    working-directory: '.'              # Working directory (default)
    test-args: '-coverprofile=coverage.out'  # Additional test args
    golangci-lint-version: 'v1.55.2'   # Specific linter version
    lint-args: '--enable-all'          # Additional lint args
    lint-timeout: '10m'                 # Lint timeout (default: 5m)
    check-tools: 'true'                 # Check tools/go.mod (default: true)
```

For detailed configuration options, see the README in each action's directory.
