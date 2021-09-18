package iset

import (
	"github.com/grobza/gasm/gasmerr"
	"github.com/grobza/gasm/parser"
)

type Constructor func([]interface{}) (Instruction, error)

// Registry instruction names registry
type Registry struct {
	data map[string]Constructor
}

func NewRegistry() *Registry {
	return &Registry{
		data: map[string]Constructor{},
	}
}

func (r *Registry) Add(name string, constructor Constructor) {
	r.data[name] = constructor
}

func (r *Registry) Has(name string) bool {
	_, ok := r.data[name]
	return ok
}

func (r *Registry) Instantiate(expr parser.Expression) (Instruction, error) {
	if !r.Has(expr.Key) {
		return nil, gasmerr.ErrUnknownInstruction
	}

	args := make([]int, 0, len(expr.Operands))
	for _, operand := range expr.Operands {
		args = append(args, *operand.Number)
	}

	instr, err := r.data[expr.Key](intsToInterfaces(args))
	return instr, err
}

func intsToInterfaces(args []int) []interface{} {
	res := make([]interface{}, 0, len(args))
	for _, arg := range args {
		res = append(res, arg)
	}
	return res
}
