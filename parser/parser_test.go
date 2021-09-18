package parser

import (
	"testing"

	"github.com/alecthomas/repr"
)

func TestParser(t *testing.T) {
	prog := &GASM{}
	err := parser.ParseString("", `
.regs 10
mov $2 4
ddd 5000
ret
`, prog)
	if err != nil {
		t.Fatal(err)
	}
	repr.Println(prog, repr.Indent("  "), repr.OmitEmpty(true))
}
