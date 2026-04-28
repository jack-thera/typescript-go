package compiler

// SyntaxKind represents the kind of a syntax node in the TypeScript AST.
type SyntaxKind int

const (
	SyntaxKindUnknown SyntaxKind = iota
	SyntaxKindEndOfFile
	SyntaxKindSingleLineComment
	SyntaxKindMultiLineComment
	SyntaxKindNewLineTrivia
	SyntaxKindWhitespaceTrivia

	// Literals
	SyntaxKindNumericLiteral
	SyntaxKindBigIntLiteral
	SyntaxKindStringLiteral
	SyntaxKindJsxText
	SyntaxKindRegularExpressionLiteral
	SyntaxKindNoSubstitutionTemplateLiteral

	// Identifiers
	SyntaxKindIdentifier
	SyntaxKindPrivateIdentifier

	// Keywords
	SyntaxKindBreakKeyword
	SyntaxKindCaseKeyword
	SyntaxKindCatchKeyword
	SyntaxKindClassKeyword
	SyntaxKindConstKeyword
	SyntaxKindContinueKeyword
	SyntaxKindDebuggerKeyword
	SyntaxKindDefaultKeyword
	SyntaxKindDeleteKeyword
	SyntaxKindDoKeyword
	SyntaxKindElseKeyword
	SyntaxKindEnumKeyword
	SyntaxKindExportKeyword
	SyntaxKindExtendsKeyword
	SyntaxKindFalseKeyword
	SyntaxKindFinallyKeyword
	SyntaxKindForKeyword
	SyntaxKindFunctionKeyword
	SyntaxKindIfKeyword
	SyntaxKindImportKeyword
	SyntaxKindInKeyword
	SyntaxKindInstanceOfKeyword
	SyntaxKindLetKeyword
	SyntaxKindNewKeyword
	SyntaxKindNullKeyword
	SyntaxKindReturnKeyword
	SyntaxKindSuperKeyword
	SyntaxKindSwitchKeyword
	SyntaxKindThisKeyword
	SyntaxKindThrowKeyword
	SyntaxKindTrueKeyword
	SyntaxKindTryKeyword
	SyntaxKindTypeOfKeyword
	SyntaxKindVarKeyword
	SyntaxKindVoidKeyword
	SyntaxKindWhileKeyword
	SyntaxKindWithKeyword

	// Declarations
	SyntaxKindFunctionDeclaration
	SyntaxKindClassDeclaration
	SyntaxKindInterfaceDeclaration
	SyntaxKindTypeAliasDeclaration
	SyntaxKindEnumDeclaration
	SyntaxKindModuleDeclaration
	SyntaxKindVariableDeclaration
	SyntaxKindVariableDeclarationList
	SyntaxKindVariableStatement

	// Expressions
	SyntaxKindCallExpression
	SyntaxKindBinaryExpression
	SyntaxKindPrefixUnaryExpression
	SyntaxKindPostfixUnaryExpression
	SyntaxKindPropertyAccessExpression
	SyntaxKindElementAccessExpression
	SyntaxKindArrowFunction
	SyntaxKindParenthesizedExpression
	SyntaxKindObjectLiteralExpression
	SyntaxKindArrayLiteralExpression

	// Statements
	SyntaxKindBlock
	SyntaxKindEmptyStatement
	SyntaxKindExpressionStatement
	SyntaxKindIfStatement
	SyntaxKindDoStatement
	SyntaxKindWhileStatement
	SyntaxKindForStatement
	SyntaxKindForInStatement
	SyntaxKindForOfStatement
	SyntaxKindReturnStatement
	SyntaxKindThrowStatement
	SyntaxKindTryStatement

	// Top-level
	SyntaxKindSourceFile
)

// NodeFlags represents flags that can be applied to AST nodes.
type NodeFlags uint32

const (
	NodeFlagsNone        NodeFlags = 0
	NodeFlagsLet         NodeFlags = 1 << iota
	NodeFlagsConst
	NodeFlagsUsing
	NodeFlagsAwaitUsing
	NodeFlagsNestedNamespace
	NodeFlagsOptionalChain
	NodeFlagsAmbient
	NodeFlagsInWithStatement
	NodeFlagsJsonFile
)

// TextRange represents a span of text in the source file.
type TextRange struct {
	Pos int
	End int
}

// Node is the base interface for all AST nodes.
type Node struct {
	Kind   SyntaxKind
	Flags  NodeFlags
	Pos    int
	End    int
	Parent *Node
}

// SourceFile represents a parsed TypeScript source file.
type SourceFile struct {
	Node
	FileName        string
	Text            string
	Statements      []*Node
	EndOfFileToken  *Node
	IsDeclarationFile bool
	HasNoDefaultLib bool
}

// Diagnostic represents a compiler error, warning, or message.
type Diagnostic struct {
	File     *SourceFile
	Start    int
	Length   int
	Code     int
	Category DiagnosticCategory
	Message  string
}

// DiagnosticCategory represents the severity of a diagnostic.
type DiagnosticCategory int

const (
	DiagnosticCategoryWarning DiagnosticCategory = iota
	DiagnosticCategoryError
	DiagnosticCategorySuggestion
	DiagnosticCategoryMessage
)

// String returns a human-readable string for the diagnostic category.
func (c DiagnosticCategory) String() string {
	switch c {
	case DiagnosticCategoryWarning:
		return "warning"
	case DiagnosticCategoryError:
		return "error"
	case DiagnosticCategorySuggestion:
		return "suggestion"
	case DiagnosticCategoryMessage:
		return "message"
	default:
		return "unknown"
	}
}
