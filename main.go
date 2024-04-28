package goscript

import (
	"github.com/breadleaf/goscript/internal/lexer"
)

func Execute(program []string) {
	l := lexer.NewLexer(program)
	l.Lex()
}
