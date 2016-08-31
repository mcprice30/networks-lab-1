package network 

import (
	"fmt"
	"net"

	"shared/go/log"
)

func OpenTCPConnection(hostName, hostPort string) (*net.TCPConn, error){
	hostAddresses, err := net.LookupHost(hostName)
	if err != nil {
		log.Fatal("Cannot find host %s : %s", hostName, err)
	}

	for i, address := range hostAddresses {
		log.Trace("Address #%d: %s:%s", i, address, hostPort)
		if remoteAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", address, hostPort)); err != nil {
			log.Warn("Error in resolution: %s", err)
		} else {
			log.Info("Found address: %s", remoteAddr)
			conn, err := net.DialTCP(remoteAddr.Network(), nil, remoteAddr)
			if err != nil {
				log.Warn("Error dialing: %s\n", err)
			}	else {
				return conn, nil
			}
		}
	}
	return nil, fmt.Errorf("Could not find host!")
}
