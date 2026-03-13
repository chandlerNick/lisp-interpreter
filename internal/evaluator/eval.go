package evaluator

import (
    "github.com/chandlernick/lisp-interpreter/internal/ast"
    "github.com/chandlernick/lisp-interpreter/internal/env"
)

func Eval(node ast.Node, e *env.Environment) interface{} {
    switch n := node.(type) {
    case *ast.IntegerNode:
        return n.Value
    case *ast.IdentifierNode:
        return e.Get(n.TokenLiteral()) // Lookup variable/func in env
    case *ast.CallNode:
        return evalCall(n, e)
    }
    return nil
}

func evalCall(node *ast.CallNode, e *env.Environment) interface{} {
    // Evaluate the function node
    fn := Eval(node.Function, e).(func([]interface{}) interface{})

    args := []interface{}{}
    for _, arg := range node.Arguments {
        args = append(args, Eval(arg, e))
    }
    return fn(args)
}
