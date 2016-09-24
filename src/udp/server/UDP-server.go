package main

import (
		"fmt"
		"net"
		"os"

		"shared/go/types"
		"shared/go/log"
)

const Port string = "10010"

func main() {

		if len(os.Args) != 3 && len(os.Args) != 2 {
			fmt.Fprintf(os.Stderr, "Usage: %s port [LOGLEVEL]\n", os.Args[0])	
			os.Exit(1)
		}

		if len(os.Args) == 3 {
			if err := log.Level(os.Args[2]); err != nil {
				log.Fatal("Invalid log level: %s", os.Args[2])
			}
		}

		serverAddr, err := net.ResolveUDPAddr("udp", ":" + os.Args[1])
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		conn, err := net.ListenUDP("udp", serverAddr)
		if err != nil {
			log.Fatal(err.Error())
			return	
		}

		defer conn.Close()

		buf := make([]byte, 1024)

		for {
			n,addr,err := conn.ReadFromUDP(buf)
			log.Trace("Recieved %+v from %s\n",	buf[0:n], addr)		

	
			if err != nil {
				log.Error(err.Error())
			}	

			request, err := types.CalcRequestFromBytes(buf[0:n])
			if err != nil {
				log.Error(err.Error())
			}

			log.Trace("Got request %+v", request)
			response := DoMath(request)
			log.Trace("Sending: %s", response.ToBytes())

			if n, err := conn.WriteTo(response.ToBytes(), addr); err != nil {
				log.Error(err.Error())
			} else {
				log.Info("Sent %d bytes", n)
			}
		}

}

func DoMath(req *types.CalculationRequest) (*types.CalculationResponse) {

	reqID := req.RequestID
	op1 := types.Result(req.Operand1)
	op2 := types.Result(req.Operand2)

	switch req.OpCode {
	case types.OpCodeAdd:
		return types.BuildResponse(reqID, op1 + op2)
	case types.OpCodeSub:
		return types.BuildResponse(reqID, op1 - op2)
	case types.OpCodeOr:
		return types.BuildResponse(reqID, op1 | op2)
	case types.OpCodeAnd:
		return types.BuildResponse(reqID, op1 & op2)
	case types.OpCodeRShift:
		return types.BuildResponse(reqID, op1 >> uint32(op2))
	case types.OpCodeLShift:
		return types.BuildResponse(reqID, op1 << uint32(op2))
	default:
		return types.ErrResponse(reqID)
	}
}
