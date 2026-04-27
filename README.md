# typescript-go

A Go implementation of the TypeScript compiler and language services.

This is a fork of [microsoft/typescript-go](https://github.com/microsoft/typescript-go), aimed at providing a high-performance TypeScript compiler written in Go.

## Overview

`typescript-go` reimplements the TypeScript compiler in Go, offering:

- **Faster compilation** — leveraging Go's concurrency and performance characteristics
- **Lower memory usage** — more efficient data structures and memory management
- **Native binaries** — no Node.js runtime required for compilation
- **Language server support** — LSP-compatible language server for editor integrations

## Status

> ⚠️ This project is under active development. APIs and behavior may change.

## Requirements

- Go 1.22 or later
- Node.js 18+ (for running TypeScript test suites)

## Getting Started

### Building

```bash
git clone https://github.com/nicholasgasior/typescript-go
cd typescript-go
go build ./...
```

### Running Tests

```bash
go test ./...
```

### Running the Compiler

```bash
go run ./cmd/tsgo --project tsconfig.json
```

## Development

### Dev Container

This repository includes a [Dev Container](.devcontainer/devcontainer.json) configuration for VS Code. Open the repository in VS Code and select **Reopen in Container** to get a fully configured development environment.

### Code Formatting

We use [dprint](https://dprint.dev/) for formatting. Configuration is in [`.dprint.jsonc`](.dprint.jsonc).

```bash
dprint fmt
```

### Linting

We use [golangci-lint](https://golangci-lint.run/) with a custom configuration defined in [`.custom-gcl.yml`](.custom-gcl.yml).

```bash
golangci-lint run
```

## Project Structure

```
typescript-go/
├── cmd/              # Entry points (compiler CLI, language server)
├── internal/         # Internal packages
│   ├── ast/          # TypeScript AST definitions
│   ├── checker/      # Type checker
│   ├── parser/       # TypeScript parser
│   ├── scanner/      # Lexer/scanner
│   └── compiler/     # Compiler pipeline
├── testdata/         # Test fixtures
└── tests/            # Integration tests
```

## Contributing

Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

When filing issues, use the appropriate [issue template](.github/ISSUE_TEMPLATE/):
- **Crash** — for panics or unexpected crashes
- **Behavior difference** — for output differences vs the official TypeScript compiler
- **VS Code editor issue** — for language server / editor integration issues
- **Other** — for anything else

## License

This project is licensed under the Apache 2.0 License — see the [LICENSE](LICENSE) file for details.

Portions of this code are derived from [microsoft/TypeScript](https://github.com/microsoft/TypeScript), which is licensed under the Apache 2.0 License.
