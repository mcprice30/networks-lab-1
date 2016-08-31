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

	reqNum := byte(0)
	for {
		reqBytes := types.BuildRequest(reqNum).ToBytes()
		reqNum++
		if n, err := conn.Write(reqBytes); err != nil	{
			log.Error("Error: %s", err)	
		} else {
			log.Info("Sent %d bytes", n)
		}
	}

} 
