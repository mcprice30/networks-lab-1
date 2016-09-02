package types 

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"

	"shared/go/log"
)

type Operand int16
type NumOperands byte
type OpCode byte

const (
	RequestTML TML = 8
	RequestNumOperands NumOperands = 2
)

type CalculationRequest struct {
	TML TML
	RequestID RequestID
	OpCode OpCode
	NumOperands NumOperands
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
		TML: TML(in[0]),
		RequestID: RequestID(in[1]),
		OpCode: OpCode(in[2]),
		NumOperands: NumOperands(in[3]),
		Operand1: op1,
		Operand2: op2,
	}, nil
}

func (req *CalculationRequest) ToBytes() []byte {
	out := make([]byte, 0, int(req.TML))
	out = append(out, byte(req.TML))
	out = append(out, byte(req.RequestID))
	out = append(out, byte(req.OpCode))
	out = append(out, byte(req.NumOperands))
	out = append(out, operand_toBytes(req.Operand1)...)
	out = append(out, operand_toBytes(req.Operand2)...)
	log.Trace("Request bytes: %+v", out)
	return out
}

func BuildRequest(reqNum RequestID) *CalculationRequest {
	request := &CalculationRequest{
		TML: RequestTML,
		RequestID: reqNum,
		NumOperands: RequestNumOperands,	
	}

	ReadInput("Enter opcode: ", &request.OpCode)
	ReadInput("Enter operand 1: ", &request.Operand1)
	ReadInput("Enter operand 2: ", &request.Operand2)

	return request
}

func ReadInput(prompt string, readInto interface{}) {

	satisfied := false

	stdin := bufio.NewReader(os.Stdin)

	for !satisfied {
		fmt.Print(prompt)
		var intVal int
		_, err := fmt.Scanf("%d", &intVal)
		if err != nil {
			fmt.Printf("Error reading input (%s)\n", err)
			stdin.ReadString('\n')	
			continue
		}

		satisfied = true

		switch readInto := readInto.(type) {
		case *OpCode:
			if intVal > 0xff {
				fmt.Printf("Input too large!")	
				satisfied = false	
			}
			*readInto = OpCode(intVal)
		case *Operand:
			if intVal > 0xffff {
				fmt.Printf("Input too large!")
				satisfied = false
			}
			*readInto = Operand(intVal)
		default:
			fmt.Fprintf(os.Stderr, "Can only read into byte or uint16 pointers. Got: %s!\n",
				reflect.TypeOf(readInto))
			os.Exit(1)
		}

	}
}
