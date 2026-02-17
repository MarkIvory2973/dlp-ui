package utils

import (
	"net"
)

func ListIPs() ([]string, error) {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return []string{}, err
	}

	var ips []string
	for _, address := range addresses {
		ipnet, ok := address.(*net.IPNet)
		if !ok {
			continue
		}

		if ipnet.IP.IsLoopback() {
			continue
		} else if ipnet.IP.To4() == nil {
			continue
		}

		ips = append(ips, ipnet.IP.String())
	}

	return ips, nil
}
