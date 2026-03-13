package ast

import "github.com/chandlernick/lisp-interpreter/internal/lexer"

type Node interface {
    TokenLiteral() string
    String() string
}

type IntegerNode struct { Token lexer.Token; Value int64 }
func (n *IntegerNode) TokenLiteral() string { return n.Token.Literal }
func (n *IntegerNode) String() string       { return n.Token.Literal }

type IdentifierNode struct { Token lexer.Token }
func (n *IdentifierNode) TokenLiteral() string { return n.Token.Literal }
func (n *IdentifierNode) String() string       { return n.Token.Literal }

// CallNode now holds the Function as a Node to allow dynamic resolution
type CallNode struct { 
    Function  Node   // The operator or function expression
    Arguments []Node // The arguments to apply
}

func (n *CallNode) TokenLiteral() string { 
    return n.Function.TokenLiteral() 
}

func (n *CallNode) String() string { 
    return "(" + n.Function.String() + " ...)" 
}