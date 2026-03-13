package lexer

import (
	"fmt"
	"strings"
	"text/scanner"
)

type TokenType int

const (
	ILLEGAL TokenType = iota
	EOF
	SYMBOL // Combined IDENT and operators
	NUMBER // Formerly INT
	LPAREN
	RPAREN
)

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

type Lexer struct {
	scanner scanner.Scanner
}

func NewLexer(input string) *Lexer {
	var s scanner.Scanner
	s.Init(strings.NewReader(input))
	// Ensure the scanner treats these characters as individual tokens
	s.Mode = scanner.ScanInts | scanner.ScanIdents
	return &Lexer{scanner: s}
}

// NextToken returns the next token in the stream, or EOF
func (l *Lexer) NextToken() (Token, error) {
	tok := l.scanner.Scan()
	pos := l.scanner.Pos()
	text := l.scanner.TokenText()

	if tok == scanner.EOF {
		return Token{Type: EOF, Literal: "", Line: pos.Line, Column: pos.Column}, nil
	}

	var tokenType TokenType
	switch tok {
	case '(':
		tokenType = LPAREN
	case ')':
		tokenType = RPAREN
	case '+', '-', '*', '/':
		// Operators are treated as symbols in this minimalist grammar
		tokenType = SYMBOL
	case scanner.Int:
		tokenType = NUMBER
	case scanner.Ident:
		tokenType = SYMBOL
	default:
		return Token{Type: ILLEGAL, Literal: text, Line: pos.Line, Column: pos.Column},
			fmt.Errorf("illegal character '%s' at line %d, column %d", text, pos.Line, pos.Column)
	}

	return Token{
		Type:    tokenType,
		Literal: text,
		Line:    pos.Line,
		Column:  pos.Column,
	}, nil
}