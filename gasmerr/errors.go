package gasmerr

type GasmError string

func (g GasmError) Error() string {
	return string(g)
}

const ErrStop = GasmError("halt execution")
const ErrUnknownInstruction = GasmError("unknown instruction")
const ErrWrongOperandType = GasmError("wrong operand type")
const ErrWrongNumberOfOperands = GasmError("wrong number of operands")
