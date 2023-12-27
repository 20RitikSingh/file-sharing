package main

import (
	"fmt"
	"net"
)

func getWLANIPAddress() (net.IP, error) {
	// Get all network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("error getting network interfaces: %v", err)
	}

	// Find the IP address of the first WLAN interface
	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0 {
			if iface.Name == "Wi-Fi" {
				// Get interface addresses
				addrs, err := iface.Addrs()
				if err != nil {
					return nil, fmt.Errorf("error getting addresses for interface %s: %v", iface.Name, err)
				}

				// Find the first non-loopback IP address
				for _, addr := range addrs {
					ipNet, ok := addr.(*net.IPNet)
					if ok && ipNet.IP.To4() != nil && !ipNet.IP.IsLoopback() {
						return ipNet.IP, nil
					}
				}
			}
		}
	}

	return nil, fmt.Errorf("no WLAN interface found or no IP address associated")
}
