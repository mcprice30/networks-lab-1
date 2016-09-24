package main

import (
	"fmt"
	"os"

	"shared/go/types"
	"shared/go/network"
	"shared/go/log"
)

func main() {

	if len(os.Args) != 3 && len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname port [LOGLEVEL]\n", os.Args[0])
		os.Exit(1)
	}

	if len(os.Args) == 4 {
		if err := log.Level(os.Args[3]); err != nil {
			log.Fatal("Invalid log level: %s", os.Args[3])
		}
	}

	hostName := os.Args[1]
	hostPort := os.Args[2]

	conn, err := network.OpenTCPConnection(hostName, hostPort)
	if err != nil {
		log.Fatal("Error: %s", err)
	}

	reqNum := types.RequestID(0)
	respBytes := make([]byte, 1000);
	for {
		reqBytes := types.BuildRequest(reqNum).ToBytes()
		reqNum++
		if n, err := conn.Write(reqBytes); err != nil	{
			log.Error("Error sending request: %s", err)	
		} else {
			log.Info("Sent %d bytes", n)
		}

		if bytesReturned, err := conn.Read(respBytes); err != nil {
			log.Error("Error getting response: %s", err)
		} else {
			log.Info("Recieved %d bytes", bytesReturned)
			response, err := types.CalcResponseFromBytes(respBytes[:bytesReturned])
			if err != nil {
				log.Error("Error converting bytes to a response: %s", err)
			} else {
				if response.ErrorCode	!= types.ErrorCodeNoError {
					fmt.Printf("Could not compute request #%d (Error Code %d)\n",
						response.RequestID, response.ErrorCode)
				} else {
					fmt.Printf("Result for request #%d: %d\n", response.RequestID,
						response.Result)
				}
			}
		}

	}

} 
