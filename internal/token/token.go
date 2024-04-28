package token

import (
	"strconv"
	"unicode"
)

type Token int

const (
	// Special tokens
	EOF Token = iota
	COMMENT

	// Literals
	literal_begin
	IDENTIFIER
	NUMBER
	STRING
	literal_end

	// Operators
	operator_begin

	// Arithmetic
	ADD
	SUBTRACT
	MULTIPLY
	DIVIDE
	REMAINDER

	// Logical
	LOGICAL_AND
	LOGICAL_OR
	LOGICAL_NOT

	// Comparison
	EQUAL
	NOT_EQUAL
	LESS_THAN
	LESS_THAN_EQUAL
	GREATER_THAN
	GREATER_THAN_EQUAL

	// Assignment
	ASSIGN

	// Delimiters
	COMMA
	SEMICOLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE

	operator_end

	// Keywords
	keyword_begin
	LET
	CONST

	IF
	ELSE

	FOR
	BREAK
	CONTINUE

	FUNCTION
	RETURN
	keyword_end
)

var tokens = [...]string{
	EOF: "EOF",
	COMMENT: "COMMENT",

	IDENTIFIER: "IDENTIFIER",
	NUMBER: "NUMBER",
	STRING: "STRING",

	ADD: "+",
	SUBTRACT: "-",
	MULTIPLY: "*",
	DIVIDE: "/",
	REMAINDER: "%",

	LOGICAL_AND: "&&",
	LOGICAL_OR: "||",
	LOGICAL_NOT: "!",
	EQUAL: "==",
	NOT_EQUAL: "!=",
	LESS_THAN: "<",
	LESS_THAN_EQUAL: "<=",
	GREATER_THAN: ">",
	GREATER_THAN_EQUAL: ">=",

	ASSIGN: "=",

	COMMA: ",",
	SEMICOLON: ";",
	LPAREN: "(",
	RPAREN: ")",
	LBRACE: "{",
	RBRACE: "}",

	LET: "let",
	CONST: "const",

	IF: "if",
	ELSE: "else",

	FOR: "for",
	BREAK: "break",
	CONTINUE: "continue",

	FUNCTION: "function",
	RETURN: "return",
}

func (t Token) String() string {
	s := ""

	if 0 <= t && t < Token(len(tokens)) {
		s = tokens[t]
	}

	if s == "" {
		s = "token(" + strconv.Itoa(int(t)) + ")"
	}

	return s
}

func (op Token) Precedence() int {
	switch op {
		case LOGICAL_OR:
			return 1
		case LOGICAL_AND:
			return 2
		case EQUAL, NOT_EQUAL, LESS_THAN, LESS_THAN_EQUAL, GREATER_THAN, GREATER_THAN_EQUAL:
			return 3
		case ADD, SUBTRACT:
			return 4
		case MULTIPLY, DIVIDE, REMAINDER:
			return 5
	}

	return 0
}

var keywords map[string]Token

func init() {
	keywords = make(map[string]Token, keyword_end - (keyword_begin + 1))

	for i := keyword_begin + 1; i < keyword_end; i++ {
		keywords[tokens[i]] = i
	}
}

func Lookup(ident string) Token {
	if tok, is_keyword := keywords[ident]; is_keyword {
		return tok
	}

	return IDENTIFIER
}

func (tok Token) IsLiteral() bool {
	return literal_begin < tok && tok < literal_end
}

func (tok Token) IsOperator() bool {
	return operator_begin < tok && tok < operator_end
}

func (tok Token) IsKeyword() bool {
	return keyword_begin < tok && tok < keyword_end
}

func IsKeyword(name string) bool {
	_, is_keyword := keywords[name]
	return is_keyword
}

func IsIdentifier(name string) bool {
	if name == "" || IsKeyword(name) {
		return false
	}

	for index, char := range name {
		if !unicode.IsLetter(char) && char != '_' && (index == 0 || !unicode.IsDigit(char)) {
			return false
		}
	}

	return true
}
