package iset

import (
	"github.com/grobza/gasm/register"
)

type Instruction interface {
	Execute(state *register.State) error
	String() string
}
