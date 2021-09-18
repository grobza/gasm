package iset

import (
	"errors"
	"fmt"
	"strings"

	"github.com/grobza/gasm/gasmerr"
	"github.com/grobza/gasm/register"
)

type mov struct {
	dst int
	val interface{}
}

func Mov(args []interface{}) (Instruction, error) {
	if len(args) != 2 {
		return nil, gasmerr.ErrWrongNumberOfOperands
	}
	instr := &mov{}

	v, ok := args[0].(int)
	if !ok {
		return nil, gasmerr.ErrWrongOperandType
	}
	instr.dst = v

	v, ok = args[1].(int)
	if !ok {
		return nil, gasmerr.ErrWrongOperandType
	}
	instr.val = v

	return instr, nil
}

func (a *mov) Execute(state *register.State) error {
	regs := state.Common()
	if a.dst >= len(regs) || a.dst < 0 {
		return errors.New("register index out of range")
	}
	regs[a.dst].Set(a.val)
	state.IncIp()
	return nil
}

func (a *mov) String() string {
	b := strings.Builder{}
	b.WriteString("mov")
	b.WriteString(fmt.Sprintf(" %d", a.dst))
	b.WriteString(fmt.Sprintf(" %d", a.val))
	return b.String()
}
