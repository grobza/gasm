package gasm

import (
	"github.com/grobza/gasm/gasmerr"
	"github.com/grobza/gasm/iset"
	"github.com/grobza/gasm/parser"
	"github.com/grobza/gasm/program"
	"github.com/grobza/gasm/register"
)

type interpreter struct {
	prog  *program.Program
	state *register.State

	instReg *iset.Registry
}

func (i *interpreter) FromSource(src string) (*program.Program, error) {
	progTree := &parser.GASM{}
	err := parser.ParseSource(src, progTree)
	if err != nil {
		return nil, err
	}

	prog := &program.Program{}

	for _, parserInstr := range progTree.Instructions {
		instr, err := i.instReg.Instantiate(*parserInstr)
		if err != nil {
			return nil, err
		}
		prog.AddInstr(instr)
	}

	return prog, nil
}

func (i *interpreter) Prog() *program.Program {
	return i.prog
}

func New(regNum int) *interpreter {
	regNum = 10
	i := new(interpreter)
	i.state = register.NewState(regNum)
	i.prog = new(program.Program)

	i.instReg = iset.NewRegistry()
	i.instReg.Add("add", iset.Add)
	i.instReg.Add("mov", iset.Mov)
	i.instReg.Add("ret", iset.Ret)
	i.instReg.Add("print", iset.Print)

	return i
}

func (i *interpreter) Process(p *program.Program) error {
	var err error
	for {
		err = p.Instrs()[i.state.Ip()].Execute(i.state)
		if err != nil {
			if err == gasmerr.ErrStop {
				return nil
			}
			return err
		}
	}
}

func (i *interpreter) RegInstr(name string, constructor iset.Constructor) {
	i.instReg.Add(name, constructor)
}
