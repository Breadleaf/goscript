package printer

import (
	"testing"
	"fmt"
)

func TestPrinter(t *testing.T) {
	p := BeginColor()
	p.AddText("Hello, ")
	p.SetFormat(UNDERLINE, NONE)
	p.SetFormat(BOLD, NONE)
	p.AddText("World!")
	fmt.Println(p.Build())
}
