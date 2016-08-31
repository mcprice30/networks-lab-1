package types 

import (
	"unsafe"

	"shared/go/input"
	"shared/go/log"
)

type CalculationRequest struct {
	TML byte
	RequestID byte
	OpCode byte
	NumOperands byte
	Operand1 uint16
	Operand2 uint16
}

func uint16_toBytes(n uint16) []byte {
	return []byte{byte(n>>8), byte(n)}
}

func (req *CalculationRequest) ToBytes() []byte {
	out := make([]byte, 0, int(req.TML))
	out = append(out, req.TML)
	out = append(out, req.RequestID)
	out = append(out, req.OpCode)
	out = append(out, req.NumOperands)
	out = append(out, uint16_toBytes(req.Operand1)...)
	out = append(out, uint16_toBytes(req.Operand2)...)
	log.Trace("Request bytes: %+v", out)
	return out
}

func isLittleEndian() bool {
	i := uint16(1)
	return (*[2]uint8)(unsafe.Pointer(&i))[0] > 0
}


func BuildRequest(reqNum byte) *CalculationRequest {
	request := &CalculationRequest{
		TML: 8,
		RequestID: reqNum,
		NumOperands: 2,	
	}

	input.ReadInput("Enter opcode: ", &request.OpCode)
	input.ReadInput("Enter operand 1: ", &request.Operand1)
	input.ReadInput("Enter operand 2: ", &request.Operand2)

	return request
}
