package main

import (
	"github.com/breadleaf/goscript/internal/lexer"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	filename := os.Args[1]

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	text := string(data)
	lines := strings.Split(text, "\n")

	l := lexer.NewLexer(lines)
	l.Lex()
}

func usage() {
	fmt.Printf("Usage: %s <filename>\n", os.Args[0])
}
