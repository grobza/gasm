package parser

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type GASM struct {
	pos          lexer.Position
	Directives   []*Directive  `@@*`
	Instructions []*Expression `@@*`
}

type Directive struct {
	pos     lexer.Position
	Key     string     `@"."@Ident `
	Operand []*Operand `(@@)*`
}

type Expression struct {
	pos      lexer.Position
	Key      string     `@Ident`
	Operands []*Operand `(@@)*`
}

type Operand struct {
	pos    lexer.Position
	Number *int `"$"?@Int|@String`
}

var parser = participle.MustBuild(&GASM{}, participle.UseLookahead(2))

func ParseSource(src string, gasm *GASM) error {
	return parser.ParseString(``, src, gasm)
}
