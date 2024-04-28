// +build windows

package printer

import (
	"fmt"
)

type Style int

const (
	NORMAL Style = 0
	BOLD Style = 1
	LIGHT Style = 2
	ITALIC Style = 3
	UNDERLINE Style = 4
)

type Color int

const (
	NONE Color = 0
	BLACK Color = 30
	RED Color = 31
	GREEN Color = 32
	YELLOW Color = 33
	BLUE Color = 34
	MAGENTA Color = 35
	CYAN Color = 36
	WHITE Color = 37
)

type Printer struct {
	Text string
}

func BeginColor() *Printer {
	return &Printer{
		Text: "",
	}
}

func (p *Printer) SetFormat(style Style, color Color) *Printer {
	p.Text = fmt.Sprintf("%s", p.Text)
	return p
}

func (p *Printer) ResetFormat() *Printer {
	return p.SetFormat(NORMAL, NONE)
}

func (p *Printer) AddText(text string) *Printer {
	p.Text = fmt.Sprintf("%s%s", p.Text, text)
	return p
}

func (p *Printer) Build() string {
	return p.Text
}

func (p *Printer) String() string {
	return p.Build()
}
