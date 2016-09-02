package types

import (
	"errors"

	"shared/go/log"
)

type TML byte
type RequestID byte
type NumOperands byte
type OpCode byte
type ErrorCode byte

type Operand int16
type Result int32


func (n Operand) toBytes()[]byte {
	return int16ToBigEndianBytes(int16(n))
}

func bytesToOperand(in []byte) Operand {
	res, err := bigEndianBytesToInt16(in)
	if err != nil {
		log.Fatal("Could not build operand: %s", err)
	}
	return Operand(res)
}

func (r Result) toBytes() []byte {
	return int32ToBigEndianBytes(int32(r))
}

func bytesToResult(in []byte) Result {
	res, err := bigEndianBytesToInt32(in)
	if err != nil {
		log.Fatal("Could not build result: %s", err)
	}
	return Result(res)
}

func bigEndianBytesToInt16(in []byte) (int16, error){
	if len(in) != 2 {
		return 0, errors.New("16 bit value needs 2 bytes")	
	}
	return int16(in[0]) << 8 + int16(in[1]), nil;
}

func int16ToBigEndianBytes(in int16) []byte {
	return []byte{byte(in>>8), byte(in)}
}

func bigEndianBytesToInt32(in []byte) (int32, error) {
	if len(in) != 4 {
		return 0, errors.New("int32 needs 4 bytes!")
	}
	return int32(in[0]) << 24 + int32(in[1]) << 16 + int32(in[2]) << 8 + int32(in[3]), nil;
}

func int32ToBigEndianBytes(in int32) []byte {
	return []byte{byte(in>>24), byte(in>>16), byte(in>>8), byte(in)}
}
