package main

import (
	"github.com/breadleaf/goscript/internal/lexer"
)

func Execute(program []string) {
	l := lexer.NewLexer(program)
	l.Lex()
}

func main() {
	program := []string{
		"var a = 1",
		"var b = 2",
		"var c = a + b",
	}

	Execute(program)
}
