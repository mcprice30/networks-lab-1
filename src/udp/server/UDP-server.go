package main

import (
		"fmt"
		"net"
		"shared/go/network"
		"shared/go/types"
)

const Port string = "10010"

func main() {
		fmt.Println("Server Here!")	

		addr, err := network.GetLocalAddress()

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("ADDRESS: %+v\n", addr)

		serverAddr, err := net.ResolveUDPAddr("udp", ":10010")
		if err != nil {
			fmt.Println(err)
			return
		}

		conn, err := net.ListenUDP("udp", serverAddr)
		if err != nil {
			fmt.Println(err)
			return	
		}

		defer conn.Close()

		buf := make([]byte, 1024)

		for {
			n,addr,err := conn.ReadFromUDP(buf)
			fmt.Printf("Recieved %+v from %s\n",	buf[0:n], addr)		

	
			if err != nil {
				fmt.Println("Error: ", err)
			}	

			request, err := types.CalcRequestFromBytes(buf[0:n])
			if err != nil {
				fmt.Println("Error: ", err)
			}

			fmt.Println(request)
			response, _ := DoMath(request)
			fmt.Println(response.ToBytes())

			if n, err := conn.WriteTo(response.ToBytes(), addr); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Sent ", n, " bytes")
			}
		}

}

func DoMath(req *types.CalculationRequest) (*types.CalculationResponse, error) {

	reqID := req.RequestID
	op1 := types.Result(req.Operand1)
	op2 := types.Result(req.Operand2)

	switch req.OpCode {
	case 0:
		return types.BuildResponse(reqID, op1 + op2), nil
	case 1:
		return types.BuildResponse(reqID, op1 - op2), nil
	case 2:
		return types.BuildResponse(reqID, op1 | op2), nil
	case 3:
		return types.BuildResponse(reqID, op1 & op2), nil
	case 4:
		return types.BuildResponse(reqID, op1 >> uint32(op2)), nil
	case 5:
		return types.BuildResponse(reqID, op1 << uint32(op2)), nil
	}

	return nil, nil
}
