package lexer

import (
	"github.com/breadleaf/goscript/internal/token"
	"fmt"
	"unicode"
)

type Lexer struct {
	Program []string
	Tokens  []token.Token
}

func NewLexer(program []string) *Lexer {
	return &Lexer {
		Program: program,
		Tokens: make([]token.Token, 0),
	}
}

func (l *Lexer) Lex() {
	for lineNumber, line := range l.Program {
		for columnNumber, char := range line {
			if unicode.IsSpace(char) {
				continue
			}

			_ = lineNumber + columnNumber
			fmt.Print(string(char))
		}
	}
}
