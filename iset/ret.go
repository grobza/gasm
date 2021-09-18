package iset

import (
	"github.com/grobza/gasm/gasmerr"
	"github.com/grobza/gasm/register"
)

type ret struct {
}

func Ret(args []interface{}) (Instruction, error) {
	if len(args) != 0 {
		return nil, gasmerr.ErrWrongNumberOfOperands
	}
	instr := &ret{}

	return instr, nil
}

func (a *ret) Execute(state *register.State) error {
	state.IncIp()
	return gasmerr.ErrStop
}

func (a *ret) String() string {
	return "ret"
}
