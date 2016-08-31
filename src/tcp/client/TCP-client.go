package main

import (
	"fmt"
	"net"
	"os"
	"reflect"
	"unsafe"
)

func main() {

	fmt.Println(isLittleEndian())

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname port\n", os.Args[0])
		os.Exit(1)
	}

	hostName := os.Args[1]
	hostPort := os.Args[2]

	conn, err := openConnection(hostName, hostPort)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	reqNum := byte(0)
	for {
		reqBytes := buildRequest(&reqNum).toBytes()
		if n, err := conn.Write(reqBytes); err != nil	{
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)	
		} else {
			fmt.Printf("Sent %d bytes\n", n)
		}
	}

} 

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

func (req *CalculationRequest) toBytes() []byte {
	out := make([]byte, 0, int(req.TML))
	out = append(out, req.TML)
	out = append(out, req.RequestID)
	out = append(out, req.OpCode)
	out = append(out, req.NumOperands)
	out = append(out, uint16_toBytes(req.Operand1)...)
	out = append(out, uint16_toBytes(req.Operand2)...)
	return out
}

func isLittleEndian() bool {
	i := uint16(1)
	return (*[2]uint8)(unsafe.Pointer(&i))[0] > 0
}

func readInput(prompt string, readInto interface{}) {

	satisfied := false

	for !satisfied {
		fmt.Print(prompt)
		var intVal uint
		_, err := fmt.Scanf("%d", &intVal)
		if err != nil {
			fmt.Printf("Error reading input (%s)\n", err)
			continue
		}

		satisfied = true

		switch readInto := readInto.(type) {
		case *byte:
			if intVal > 0xff {
				fmt.Printf("Input too large!")	
				satisfied = false	
			}
			*readInto = byte(intVal)
		case *uint16:
			if intVal > 0xffff {
				fmt.Printf("Input too large!")
				satisfied = false
			}
			*readInto = uint16(intVal)
		default:
			fmt.Fprintf(os.Stderr, "Can only read into byte or uint16 pointers. Got: %s!\n",
				reflect.TypeOf(readInto))
			os.Exit(1)
		}

	}
}

func buildRequest(reqNum *byte) *CalculationRequest {
	request := &CalculationRequest{
		TML: 8,
		RequestID: *reqNum,
		NumOperands: 2,	
	}
	*reqNum = *reqNum + 1	

	readInput("Enter opcode: ", &request.OpCode)
	readInput("Enter operand 1: ", &request.Operand1)
	readInput("Enter operand 2: ", &request.Operand2)

	return request
}

func openConnection(hostName, hostPort string) (*net.TCPConn, error){
	hostAddresses, err := net.LookupHost(hostName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot find host %s : %s", hostName, err)
		os.Exit(1)
	}

	for i, address := range hostAddresses {
		fmt.Printf("Address #%d: %s:%s\n", i, address, hostPort)
		if remoteAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", address, hostPort)); err != nil {
			fmt.Fprintf(os.Stderr, "Error in resolution: %s", err)
		} else {
			fmt.Printf("Found address: %s\n", remoteAddr)
			conn, err := net.DialTCP(remoteAddr.Network(), nil, remoteAddr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error dialing: %s\n", err)
			}	else {
				return conn, nil
			}
		}
	}
	return nil, fmt.Errorf("Could not find error!")
}
