package program

import (
	"strings"

	"github.com/grobza/gasm/iset"
)

type Program struct {
	instrs []iset.Instruction
}

func (p *Program) Instrs() []iset.Instruction {
	return p.instrs
}

func (p *Program) AddInstr(instr iset.Instruction) {
	p.instrs = append(p.instrs, instr)
}

func (p *Program) Serialize() string {
	builder := strings.Builder{}
	for _, i := range p.instrs {
		builder.WriteString(i.String())
		builder.WriteString("\n")
	}
	return builder.String()
}
