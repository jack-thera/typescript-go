// Package compiler provides the core TypeScript compilation pipeline,
// including parsing, type checking, and code emission.
package compiler

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// SourceFile represents a parsed TypeScript source file.
type SourceFile struct {
	FileName    string
	Text        string
	Diagnostics []Diagnostic
	Statements  []Node
}

// DiagnosticCategory represents the severity of a diagnostic message.
type DiagnosticCategory int

const (
	DiagnosticCategoryWarning DiagnosticCategory = iota
	DiagnosticCategoryError
	DiagnosticCategoryMessage
	DiagnosticCategorySuggestion
)

// Diagnostic represents a compiler diagnostic message with location info.
type Diagnostic struct {
	File     *SourceFile
	Start    int
	Length   int
	Message  string
	Category DiagnosticCategory
	Code     int
}

// String returns a human-readable representation of the diagnostic.
func (d *Diagnostic) String() string {
	if d.File == nil {
		return fmt.Sprintf("error TS%d: %s", d.Code, d.Message)
	}
	line, col := d.File.LineAndCharOf(d.Start)
	return fmt.Sprintf("%s(%d,%d): error TS%d: %s",
		d.File.FileName, line, col, d.Code, d.Message)
}

// LineAndCharOf returns the 1-based line and character position for a given offset.
func (sf *SourceFile) LineAndCharOf(offset int) (line, char int) {
	line = 1
	char = 1
	for i, ch := range sf.Text {
		if i >= offset {
			break
		}
		if ch == '\n' {
			line++
			char = 1
		} else {
			char++
		}
	}
	return line, char
}

// Node represents a generic AST node.
type Node struct {
	Kind  NodeKind
	Start int
	End   int
	Text  string
}

// NodeKind enumerates the kinds of AST nodes.
type NodeKind int

const (
	NodeKindUnknown NodeKind = iota
	NodeKindSourceFile
	NodeKindImportDeclaration
	NodeKindExportDeclaration
	NodeKindFunctionDeclaration
	NodeKindClassDeclaration
	NodeKindInterfaceDeclaration
	NodeKindTypeAliasDeclaration
	NodeKindVariableStatement
)

// ParseFile reads and performs a lightweight parse of a TypeScript source file.
// It returns a SourceFile with basic structural information and any read errors.
func ParseFile(fileName string) (*SourceFile, error) {
	abs, err := filepath.Abs(fileName)
	if err != nil {
		return nil, fmt.Errorf("resolving path %q: %w", fileName, err)
	}

	data, err := os.ReadFile(abs)
	if err != nil {
		return nil, fmt.Errorf("reading file %q: %w", abs, err)
	}

	sf := &SourceFile{
		FileName: abs,
		Text:     string(data),
	}

	// Validate file extension.
	ext := strings.ToLower(filepath.Ext(abs))
	switch ext {
	case ".ts", ".tsx", ".d.ts":
		// supported
	default:
		sf.Diagnostics = append(sf.Diagnostics, Diagnostic{
			File:     sf,
			Message:  fmt.Sprintf("unsupported file extension %q", ext),
			Category: DiagnosticCategoryWarning,
			Code:     6054,
		})
	}

	return sf, nil
}
