package iset

import (
	"errors"
	"fmt"
	"strings"

	"github.com/grobza/gasm/gasmerr"
	"github.com/grobza/gasm/register"
)

type add struct {
	src int
	dst int
}

func Add(args []interface{}) (Instruction, error) {
	if len(args) != 2 {
		return nil, gasmerr.ErrWrongNumberOfOperands
	}
	instr := &add{}

	v, ok := args[0].(int)
	if !ok {
		return nil, gasmerr.ErrWrongOperandType
	}
	instr.dst = v

	v, ok = args[1].(int)
	if !ok {
		return nil, gasmerr.ErrWrongOperandType
	}
	instr.src = v

	return instr, nil
}

func (a *add) Execute(state *register.State) error {
	regs := state.Common()
	if a.src >= len(regs) || a.dst >= len(regs) {
		return errors.New("register index out of range")
	}
	leftOp := regs[a.src].Get().(int)
	rightOp := regs[a.dst].Get().(int)
	regs[a.dst].Set(leftOp + rightOp)
	state.IncIp()
	return nil
}

func (a *add) String() string {
	b := strings.Builder{}
	b.WriteString("add")
	b.WriteString(fmt.Sprintf(" %d", a.dst))
	b.WriteString(fmt.Sprintf(" %d", a.src))
	return b.String()
}
