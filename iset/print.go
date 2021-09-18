package iset

import (
	"errors"
	"fmt"
	"strings"

	"github.com/grobza/gasm/gasmerr"
	"github.com/grobza/gasm/register"
)

type print struct {
	src int
}

func Print(args []interface{}) (Instruction, error) {
	if len(args) != 1 {
		return nil, gasmerr.ErrWrongNumberOfOperands
	}
	instr := &print{}

	v, ok := args[0].(int)
	if !ok {
		return nil, gasmerr.ErrWrongOperandType
	}
	instr.src = v

	return instr, nil
}

func (p *print) Execute(state *register.State) error {
	regs := state.Common()
	if p.src >= len(regs) || p.src < 0 {
		return errors.New("register index out of range")
	}
	fmt.Printf("%v\n", regs[p.src])
	state.IncIp()
	return nil
}

func (p *print) String() string {
	b := strings.Builder{}
	b.WriteString("print")
	b.WriteString(fmt.Sprintf(" %d", p.src))
	return b.String()
}
