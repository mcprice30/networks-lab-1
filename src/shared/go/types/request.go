// Lab 1 - Group 11
// 
// The types package contains custom structs for that encapsulate the packets
// to be sent, including the various fields in said packets.
//
// request.go encapsulates request packets, containing various packet fields
// and various functions for converting it to and from network byte order.
package types 

import (
	"bufio"
	"fmt"
	"os"
	"reflect"

	"shared/go/log"
)


const (
	RequestTML TML = 8
	RequestNumOperands NumOperands = 2
)

// CalculationRequest contains all of the data sent in a request packet.
type CalculationRequest struct {
	TML         TML         // Total Message Length
	RequestID   RequestID   // Unique for each request of the client.
	OpCode      OpCode      // Indicates which operation is being requested.
	NumOperands NumOperands // Will always be 2.
	Operand1    Operand
	Operand2    Operand
}

// CalcRequestFromBytes takes a slice of bytes and produces a CalcuationRequest
// from them.
func CalcRequestFromBytes(in []byte) (*CalculationRequest, error) {
	if len(in) != int(RequestTML) {
		return nil, fmt.Errorf("Calc request needs %d bytes!", int(RequestTML))
	}

	op1 := bytesToOperand(in[4:6])
	op2 := bytesToOperand(in[6:8])	

	return &CalculationRequest{
		TML: TML(in[0]),
		RequestID: RequestID(in[1]),
		OpCode: OpCode(in[2]),
		NumOperands: NumOperands(in[3]),
		Operand1: op1,
		Operand2: op2,
	}, nil
}

// ToBytes converts a request to a slice of bytes in network byte order.
func (req *CalculationRequest) ToBytes() []byte {
	out := make([]byte, 0, int(req.TML))
	out = append(out, byte(req.TML))
	out = append(out, byte(req.RequestID))
	out = append(out, byte(req.OpCode))
	out = append(out, byte(req.NumOperands))
	out = append(out, req.Operand1.toBytes()...)
	out = append(out, req.Operand2.toBytes()...)
	log.Trace("Request bytes: %+v", out)
	return out
}

// Prompt the user for input and build a calculation request from it.
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

// Prompt the user for input text and write it into the "readInto" value.
// Responsible retrying the prompt until a valid input is provided.
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
