package network 

import (
	"fmt"
	"net"
	"os"

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


func GetLocalAddress() (net.Addr, error){

	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return nil, err
		}

		for _, addr := range  addrs {
			switch addr.(type) {
			case *net.IPAddr:
				return addr, nil
			}
		}
	}

	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	addrs, err := net.LookupHost(hostname)
	if err != nil {
		return nil, err
	}

	return net.ResolveIPAddr("ip", addrs[0])
}
