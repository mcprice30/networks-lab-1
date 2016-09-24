package main

import (
	"fmt"
	"os"

	"shared/go/types"
	"shared/go/network"
	"shared/go/log"
)

// Main entry point for the TCP client.
func main() {

	// Primitive argument parsing. We have 2 mandatory and 1 optional argument.
	if len(os.Args) != 3 && len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname port [LOGLEVEL]\n", os.Args[0])
		os.Exit(1)
	}

	// If the optional argument is passed, then handle it appropriately.
	if len(os.Args) == 4 {
		if err := log.Level(os.Args[3]); err != nil {
			log.Fatal("Invalid log level: %s", os.Args[3])
		}
	}

	// Get out the 2 required arguments.
	hostName := os.Args[1]
	hostPort := os.Args[2]

	// Attempt to open a connection to the specified host name/port.
	conn, err := network.OpenTCPConnection(hostName, hostPort)
	if err != nil {
		log.Fatal("Error: %s", err)
	}

	// Initialize variables used at each request.
	reqNum := types.RequestID(0) 			// Tracks the current request ID.	 
	respBytes := make([]byte, 1000)		// Buffer to hold response from server.

	// Loop for the duration of the program.
	for {

		// Build a request based on user input, produce bytes for it.
		reqBytes := types.BuildRequest(reqNum).ToBytes()
		reqNum++

		// Send the request to the server.
		if n, err := conn.Write(reqBytes); err != nil	{
			log.Error("Error sending request: %s", err)	
		} else {
			log.Info("Sent %d bytes", n)
		}

		// Read back the response. 
		if bytesReturned, err := conn.Read(respBytes); err != nil {
			log.Error("Error getting response: %s", err)
		} else { // If we didn't get a networking error:

			// Convert the bytes recieved into a response type.
			log.Info("Recieved %d bytes", bytesReturned)
			response, err := types.CalcResponseFromBytes(respBytes[:bytesReturned])

			if err != nil {
				log.Error("Error converting bytes to a response: %s", err)
			} else {

				// Print the response if the conversion was successful.
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
