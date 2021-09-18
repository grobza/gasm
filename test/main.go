package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/grobza/gasm"
	"github.com/grobza/gasm/gasmerr"
	"github.com/grobza/gasm/iset"
	"github.com/grobza/gasm/register"
)

func main() {
	i := gasm.New(2)
	i.RegInstr("pow", Pow)
	p, err := i.FromSource(`
	mov 0 3
	mov 1 4
	add 0 1
	pow 1 2 4
	add 0 1
	print 0
	ret
`)
	if err != nil {
		log.Fatal(err)
	}
	err = i.Process(p)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Print(p.Serialize())
}

type pow struct {
	dst int
	a   int
	b   int
}

func Pow(args []interface{}) (iset.Instruction, error) {
	if len(args) != 3 {
		return nil, gasmerr.ErrWrongNumberOfOperands
	}
	instr := &pow{}

	v, ok := args[0].(int)
	if !ok {
		return nil, gasmerr.ErrWrongOperandType
	}
	instr.dst = v

	v, ok = args[1].(int)
	if !ok {
		return nil, gasmerr.ErrWrongOperandType
	}
	instr.a = v

	v, ok = args[2].(int)
	if !ok {
		return nil, gasmerr.ErrWrongOperandType
	}
	instr.b = v

	return instr, nil
}

func (a *pow) Execute(state *register.State) error {
	regs := state.Common()
	if a.dst >= len(regs) || a.dst < 0 {
		return errors.New("register index out of range")
	}
	regs[a.dst].Set(int(math.Pow(float64(a.a), float64(a.b))))
	state.IncIp()
	return nil
}

func (a *pow) String() string {
	b := strings.Builder{}
	b.WriteString("add")
	b.WriteString(fmt.Sprintf(" %d", a.dst))
	b.WriteString(fmt.Sprintf(" %d", a.a))
	b.WriteString(fmt.Sprintf(" %d", a.b))
	return b.String()
}
