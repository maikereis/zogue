package ast

import (
	"bytes"
	"zogue/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

// Represents **statements** (e.g., `let x = 5`).
// Embeds `Node` (inherits `TokenLiteral()`)
// `statementNode()` is a **marker method** -- it does nothing but lets the compiler distinguish statements from expresssions.
type Statement interface {
	Node
	statementNode() // Unexported method
}

// Represents **expressions** (e.g., `x + 5`, `"hello"`)
// Similar to `Statement`, uses a marker method `expressionNode()`
type Expression interface {
	Node
	expressionNode() // Unexported method
}

// This is the root of the AST.
// A program is a list of statements.
type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// Returns the literal value of the first token in the program (typically for debugging or printing).
// If no statements exist, return an empty string.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Represents a `let` statement like:
// let x = 5;
// `Token`: stores the actual token.Token for the let keyword.
// `Name`: the identifier being declared (e.g., `x`).
// `Value`: the expression being assigned (e.g., `5`)
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// `statementNode()` marks this as a statement.
// It is a marker method used to satisfy the Statement interface.
// It doesn't do anything at runtime but allows the compiler to distinguish
// LetStatement from expressions.
func (ls *LetStatement) statementNode() {}

// Returns the literal value of the token associated with this statement.
// Tipically, this would be "let" for a LetStatement.
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type Identifier struct {
	Token token.Token // The token.IDENT token, which holds the literal identifier name.
	Value string      // The name of the identifier (e.g., "x", "y").
}

// `expressionNode()` marks this as an expresson.
// Like `statementNode()`, it's a marker used to satisfy the Expression interface.
func (i *Identifier) expressionNode() {}

// Returns the literal value of the code associated with the identifier.
// For example, if the source code has `x`, this returns "x".
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) String() string { return i.Value }

/*
let x = y;


Program{
	Statements: []Statement{
		&LetStatement{
			Token: token.Token{Type: LET, Literal: "let"},
			Name: &Identifier{
				Token: token.Token{Type: IDENT, Literal: "x"},
				Value: "x",
			},
			Value: &Identifier{
				Token: token.Token{Type: IDENT, Literal: "y"},
				Value: "y",
			},
		},
	},
}
*/

// Represents a `return` statement like:
// return 5;
// or
// return add(15);
// `Token`: stores the actual token.Token for the return keyword.
// `ReturnValue`: the expresssion being returned (e.g., `5`, `add(15)`in neovim how can i jump to next Uppercase consoant?)
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

// `statementNode()` marks this as an statement.
func (rs *ReturnStatement) statementNode() {}

// Returns the literal value of the token associated with this statement.
// It will be the "return".
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// Represents a **expression statement** (e.g., `x+10`)
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
