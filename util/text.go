//go:generate go run generate/main.go
package util

import (
	"fmt"
	"os"

	"github.com/andrew-d/go-termutil"
)

var IsTTY bool

func init() {
	IsTTY = termutil.Isatty(os.Stdout.Fd())
}

const end = "\033[0m"

type Text struct {
	text   string
	colour string
}

func (t Text) String() string {
	if t.colour == "" || !IsTTY {
		return t.text
	}

	return fmt.Sprintf("%s%s%s", t.colour, t.text, end)
}

func (t Text) Len() int {
	return len(t.text)
}

func Plain(text string) Text {
	return Text{
		text:   text,
		colour: "",
	}
}

func ClearScreen() {
	if IsTTY {
		fmt.Print("\033[1;1H\033[2J")
	} else {
		fmt.Println()
	}
}

func ClearLine() {
	if IsTTY {
		fmt.Print("\033[1G\033[2K")
	} else {
		fmt.Println()
	}
}