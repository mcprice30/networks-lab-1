package types 

import (
	"errors"
	"unsafe"

	"shared/go/input"
	"shared/go/log"
)

type Operand int16

type CalculationRequest struct {
	TML byte
	RequestID byte
	OpCode byte
	NumOperands byte
	Operand1 Operand
	Operand2 Operand
}

func operand_toBytes(n Operand) []byte {
	return []byte{byte(n>>8), byte(n)}
}

func bytes_toOperand(n []byte) (Operand, error) {
	if len(n) != 2 {
		return 0, errors.New("uint16 needs 2 bytes!")
	}
	return Operand(n[0]) << 8 + Operand(n[1]), nil;
}

func CalcRequestFromBytes(in []byte) (*CalculationRequest, error) {
	if len(in) != 8 {
		return nil, errors.New("Calc request needs 8 bytes!")
	}

	var op1, op2 Operand
	var err error

	if op1, err = bytes_toOperand(in[4:6]); err != nil {
		return nil, err
	}

	if op2, err = bytes_toOperand(in[6:8]); err != nil {
		return nil, err
	}

	return &CalculationRequest{
		TML: in[0],
		RequestID: in[1],
		OpCode: in[2],
		NumOperands: in[3],
		Operand1: op1,
		Operand2: op2,
	}, nil
}

func (req *CalculationRequest) ToBytes() []byte {
	out := make([]byte, 0, int(req.TML))
	out = append(out, req.TML)
	out = append(out, req.RequestID)
	out = append(out, req.OpCode)
	out = append(out, req.NumOperands)
	out = append(out, operand_toBytes(req.Operand1)...)
	out = append(out, operand_toBytes(req.Operand2)...)
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
