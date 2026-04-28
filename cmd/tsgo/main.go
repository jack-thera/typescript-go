// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// tsgo is a Go implementation of the TypeScript compiler.
package main

import (
	"fmt"
	"os"

	"github.com/microsoft/typescript-go/internal/compiler"
	"github.com/microsoft/typescript-go/internal/diagnostics"
)

const version = "0.1.0"

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		printUsage()
		return nil
	}

	switch args[0] {
	case "--version", "-v":
		fmt.Printf("tsgo version %s\n", version)
		return nil
	case "--help", "-h":
		printUsage()
		return nil
	default:
		return compile(args)
	}
}

func compile(args []string) error {
	opts, files, err := parseArgs(args)
	if err != nil {
		return err
	}

	c, err := compiler.New(opts)
	if err != nil {
		return fmt.Errorf("failed to create compiler: %w", err)
	}

	result := c.Compile(files)
	for _, diag := range result.Diagnostics {
		printDiagnostic(diag)
	}

	if result.HasErrors() {
		return fmt.Errorf("compilation failed with %d error(s)", result.ErrorCount())
	}

	// Print a success message when compilation completes without errors.
	fmt.Fprintf(os.Stderr, "Compiled %d file(s) successfully.\n", len(files))

	return nil
}

func parseArgs(args []string) (*compiler.Options, []string, error) {
	// Default to strict mode for safer type checking out of the box.
	opts := compiler.DefaultOptions()
	opts.Strict = true
	var files []string

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch {
		case arg == "--strict":
			opts.Strict = true
		case arg == "--no-strict":
			opts.Strict = false
		case arg == "--noEmit":
			opts.NoEmit = true
		case arg == "--declaration":
			opts.Declaration = true
		case arg == "--outDir":
			if i+1 >= len(args) {
				return nil, nil, fmt.Errorf("--outDir requires a value")
			}
			i++
			opts.OutDir = args[i]
		case arg == "--target":
			if i+1 >= len(args) {
				return nil, nil, fmt.Errorf("--target requires a value")
			}
			i++
			opts.Target = args[i]
		case len(arg) > 0 && arg[0] != '-':
			files = append(files, arg)
		default:
			return nil, nil, fmt.Errorf("unknown option: %s", arg)
		}
	}

	if len(files) == 0 {
		return nil, nil, fmt.Errorf("no input files specified")
	}

	return opts, files, nil
}

func printDiagnostic(diag diagnostics.Diagnostic) {
	if diag.File != "" {
		fmt.Fprintf(os.Stderr, "%s(%d,%d): %s TS%d: %s\n",
			diag.File, diag.Line, diag.Character,
			diag.Category, diag.Code, diag.Message)
	} else {
		fmt.Fprintf(os.Stderr, "%s TS%d: %s\n",
			diag.Category, diag.Code, diag.Message)
	}
}

func printUsage() {
	fmt.Printf(`tsgo - TypeScript compiler (Go implementation) v%s

Usage:
  tsgo [options] <files...>

Options:
  --strict         Enable strict type checking (default: on)
  --no-strict      Disable strict type checking
  --noEmit         Do not emit output files
  --declaration    Generate .d.ts declaration files
  --outDir <dir>   Redirect output to directory
  --target <es>    Set ECMAScript target version (e.g. ES2020)
  --version        Print version
  --help           Print this help message
`, version)
}
