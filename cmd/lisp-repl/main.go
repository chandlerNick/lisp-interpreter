package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/chandlernick/lisp-interpreter/internal/env"
	"github.com/chandlernick/lisp-interpreter/internal/evaluator"
	"github.com/chandlernick/lisp-interpreter/internal/lexer"
	"github.com/chandlernick/lisp-interpreter/internal/parser"
)

func main() {
	fmt.Println("Welcome to the Lisp REPL!")
	globalEnv := env.NewEnvironment()
	scanner := bufio.NewScanner(os.Stdin)

	for { // Go's 'while true'
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		if input == "exit" {
			fmt.Println("Goodbye!")
			os.Exit(0)
		}

		// 1. Create the new live Lexer
        l := lexer.NewLexer(input)

        // 2. Pass the lexer object to the Parser
        tree, err := parser.Parse(l)
        if err != nil {
            fmt.Printf("Syntax Error: %s\n", err)
            continue
        }

        // 3. Evaluate the resulting AST
        result := evaluator.Eval(tree, globalEnv)
        fmt.Println(result)
	}
}