// Lab 1 - Group 11
// 
// The types package contains custom structs for that encapsulate the packets
// to be sent, including the various fields in said packets.
//
// base.go defines the types that various packet fields will have, along with
// providing functions for conversion to network byte order.

package types

import (
	"errors"

	"shared/go/log"
)

// Total Message Length. Prefaces both request and response packets.
type TML byte

// A unique identifier for each request. Sent in request and response packets.
type RequestID byte

// The number of operands in a given request.
type NumOperands byte

// An id that uniquely identifies the operation to perform. Sent in requests.
type OpCode byte

// An id that is 0 iff no error occurred.
type ErrorCode byte

const (
	OpCodeAdd    OpCode = 0
	OpCodeSub    OpCode = 1
	OpCodeOr	   OpCode = 2
	OpCodeAnd    OpCode = 3
	OpCodeRShift OpCode = 4
	OpCodeLShift OpCode = 5
)

// Operands are sent in the request, which represents <operand> <op> <operand>
type Operand int16

// Results are sent in the response, and is the result of the calculation.
type Result int32

// Convert an operand to bytes.
func (n Operand) toBytes()[]byte {
	return int16ToBigEndianBytes(int16(n))
}

// Convert a pair of bytes to an Oprerand type.
func bytesToOperand(in []byte) Operand {
	res, err := bigEndianBytesToInt16(in)
	if err != nil {
		log.Fatal("Could not build operand: %s", err)
	}
	return Operand(res)
}

// Convert a result to a byte slice.
func (r Result) toBytes() []byte {
	return int32ToBigEndianBytes(int32(r))
}

// Convert a 4 byte slice to a Result type.
func bytesToResult(in []byte) Result {
	res, err := bigEndianBytesToInt32(in)
	if err != nil {
		log.Fatal("Could not build result: %s", err)
	}
	return Result(res)
}

// Convert network order bytes to an int16.
func bigEndianBytesToInt16(in []byte) (int16, error){
	if len(in) != 2 {
		return 0, errors.New("16 bit value needs 2 bytes")	
	}
	return int16(in[0]) << 8 + int16(in[1]), nil;
}

// Convert an int16 to a pair of bytes in network byte order..
func int16ToBigEndianBytes(in int16) []byte {
	return []byte{byte(in>>8), byte(in)}
}

// Convert network order bytes to an int32.
func bigEndianBytesToInt32(in []byte) (int32, error) {
	if len(in) != 4 {
		return 0, errors.New("int32 needs 4 bytes!")
	}
	return int32(in[0]) << 24 + int32(in[1]) << 16 + int32(in[2]) << 8 + int32(in[3]), nil;
}

// Convert an int32 to network order bytes.
func int32ToBigEndianBytes(in int32) []byte {
	return []byte{byte(in>>24), byte(in>>16), byte(in>>8), byte(in)}
}
