package main

import (
		"fmt"
		"net"
		"os"

		"shared/go/types"
		"shared/go/log"
)

func main() {

		// Basic argument parsing. Accept a port and an optional log level.
		if len(os.Args) != 3 && len(os.Args) != 2 {
			fmt.Fprintf(os.Stderr, "Usage: %s port [LOGLEVEL]\n", os.Args[0])	
			os.Exit(1)
		}

		// If we recieved a log level argument, set the log level.
		if len(os.Args) == 3 {
			if err := log.Level(os.Args[2]); err != nil {
				log.Fatal("Invalid log level: %s", os.Args[2])
			}
		}

		// Get the specified port.
		serverAddr, err := net.ResolveUDPAddr("udp", ":" + os.Args[1])
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		// Start listening on the given port.
		conn, err := net.ListenUDP("udp", serverAddr)
		if err != nil {
			log.Fatal(err.Error())
			return	
		}

		// When the server goes down, kill the connection.
		defer conn.Close()

		// Allocate buffer for holding requests.
		buf := make([]byte, 1024)

		for {

			// Get a request.
			n, addr, err := conn.ReadFromUDP(buf)
			log.Trace("Recieved %+v from %s\n",	buf[0:n], addr)		

	
			if err != nil {
				log.Error(err.Error())
			}	

			// Convert the request bytes to a request struct.
			request, err := types.CalcRequestFromBytes(buf[0:n])
			if err != nil {
				log.Error(err.Error())
			}

			log.Trace("Got request %+v", request)
			response := doMath(request) // Build the response.
			log.Trace("Sending: %s", response.ToBytes())

			// Send the response to the client.
			if n, err := conn.WriteTo(response.ToBytes(), addr); err != nil {
				log.Error(err.Error())
			} else {
				log.Info("Sent %d bytes", n)
			}
		}
}

// doMath takes a request packet and produces the appropriate respose packet.
func doMath(req *types.CalculationRequest) (*types.CalculationResponse) {

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
