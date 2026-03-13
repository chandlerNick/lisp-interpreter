package parser

import (
	"fmt"
	"strconv"
	"github.com/chandlernick/lisp-interpreter/internal/ast"
	"github.com/chandlernick/lisp-interpreter/internal/lexer"
)


type Parser struct {
	lexer *lexer.Lexer
	curr lexer.Token
	peek lexer.Token
}

func (p *Parser) nextToken() error {
    p.curr = p.peek
    var err error
    p.peek, err = p.lexer.NextToken()
    return err
}

func (p *Parser) Parse() (ast.Node, error) {
    // Initial priming of the pump
    if err := p.nextToken(); err != nil { return nil, err }
    if err := p.nextToken(); err != nil { return nil, err }
    return ParseExpression(p)
}

func ParseExpression(p *Parser) (ast.Node, error) {
	switch p.curr.Type {
	case lexer.LPAREN:
		return ParseList(p)
	case lexer.NUMBER:
		return ParseInteger(p)
	case lexer.SYMBOL:
		return ParseIdentifier(p)
	default:
		return nil, fmt.Errorf("unexpected token: %s at line %d", p.curr.Literal, p.curr.Line)
	}
}

func ParseList(p *Parser) (ast.Node, error) {
    p.nextToken() // consume '('
    
    // First element is the operator/function
    fn, err := ParseExpression(p)
    if err != nil { return nil, err }
    
    list := &ast.CallNode{Function: fn}
    p.nextToken()

    // Subsequent elements are arguments
    for p.curr.Type != lexer.RPAREN && p.curr.Type != lexer.EOF {
        expr, err := ParseExpression(p)
        if err != nil { return nil, err }
        list.Arguments = append(list.Arguments, expr)
        p.nextToken()
    }
    return list, nil
}

func ParseInteger(p *Parser) (ast.Node, error) {
	val, _ := strconv.ParseInt(p.curr.Literal, 10, 64)
	return &ast.IntegerNode{Token: p.curr, Value: val}, nil
}

func ParseIdentifier(p *Parser) (ast.Node, error) {
	return &ast.IdentifierNode{Token: p.curr}, nil
}

// Parse converts tokens into an AST
func Parse(lex *lexer.Lexer) (ast.Node, error) {
    p := &Parser{lexer: lex}
    return p.Parse() // This calls the method you defined above!
}