# GoReleaser Release Action

Runs GoReleaser to create a release.

## Usage

```yaml
- name: Run GoReleaser release
  uses: sivchari/actions-kit/goreleaser-release@main
  with:
    go-version-file: 'go.mod'  # optional, default: 'go.mod'
    working-directory: '.'     # optional, default: '.'
    goreleaser-version: '~> v2'  # optional, default: '~> v2'
    args: '--clean'            # optional, default: '--clean'
    github-token: ${{ secrets.GITHUB_TOKEN }}  # optional
```

## Inputs

| Input | Description | Required | Default |
|-------|-------------|----------|---------|
| `go-version-file` | Path to go.mod file for Go version detection | No | `go.mod` |
| `working-directory` | Working directory to run GoReleaser in | No | `.` |
| `goreleaser-version` | Version of GoReleaser to use | No | `~> v2` |
| `args` | Additional arguments to pass to GoReleaser release | No | `--clean` |
| `github-token` | GitHub token for creating releases | No | `${{ github.token }}` |

## Example

```yaml
name: Release
on:
  push:
    tags:
      - "v*"

permissions:
  contents: write

jobs:
  release:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0
      - name: Run GoReleaser
        uses: sivchari/actions-kit/goreleaser-release@main
        env:
          HOMEBREW_TAP_GITHUB_TOKEN: ${{ secrets.HOMEBREW_TAP_GITHUB_TOKEN }}
```