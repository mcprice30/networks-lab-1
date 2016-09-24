// Lab 1 - Group 11
// 
// The network package includes utility functions for network requests, to
// remove complexity from the main entry point.
//
package network 

import (
	"fmt"
	"net"
	"os"

	"shared/go/log"
)

// OpenTCPConnection will attempt to open a TCP connection at the given host
// name and port. It will return a pointer to a TCP connection, or an error
// if it was unable to open the connection. 
func OpenTCPConnection(hostName, hostPort string) (*net.TCPConn, error){
	hostAddresses, err := net.LookupHost(hostName)
	if err != nil {
		log.Fatal("Cannot find host %s : %s", hostName, err)
	}

	for i, address := range hostAddresses {

		// Try connecting to an address of the target host.
		log.Trace("Address #%d: %s:%s", i, address, hostPort)

		// Attempt to resolve the given address.	
		if remoteAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s",
				address, hostPort)); err != nil {

			log.Warn("Error in resolution: %s", err)
		} else {

			// If resolution works, try dialing that address.
			log.Info("Found address: %s", remoteAddr)
			conn, err := net.DialTCP(remoteAddr.Network(), nil, remoteAddr) 
			if err != nil {
				log.Warn("Error dialing: %s\n", err)
			}	else {
				return conn, nil
			}
		}
	}

	// None of the host's addresses was reachable.
	return nil, fmt.Errorf("Could not find host!")
}

// GetLocalAddress() attempts to find an address that refers to this machine.
func GetLocalAddress() (net.Addr, error) {

	// Get all network interfaces.
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	// For each network interface:
	for _, i := range ifaces {

		// Get all addresses at that interface.
		addrs, err := i.Addrs()
		if err != nil {
			return nil, err
		}

		for _, addr := range  addrs {
			switch addr.(type) {
			case *net.IPAddr:		// Return the first IP address in the interface.
				return addr, nil
			}
		}
	}

	// Otherwise, if we would not find an IP address:
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	// Lookup our own hostname.
	addrs, err := net.LookupHost(hostname)
	if err != nil {
		return nil, err
	}

	// And try to resolve that to an IP address.
	return net.ResolveIPAddr("ip", addrs[0])
}
