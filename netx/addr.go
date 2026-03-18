package netx

import (
	"net"
)

// ConnectableAddr returns a new net.Addr with unspecified IPs (0.0.0.0, [::]) replaced
// by the corresponding loopback address (127.0.0.1, [::1]).
// Non-IP addresses (e.g. Unix sockets) are returned as-is.
func ConnectableAddr(addr net.Addr) net.Addr {
	switch a := addr.(type) {
	case *net.TCPAddr:
		if a.IP.IsUnspecified() {
			v := *a
			v.IP = LoopbackIP(v.IP)
			return &v
		}
	case *net.UDPAddr:
		if a.IP.IsUnspecified() {
			v := *a
			v.IP = LoopbackIP(v.IP)
			return &v
		}
	case *net.IPAddr:
		if a.IP.IsUnspecified() {
			v := *a
			v.IP = LoopbackIP(v.IP)
			return &v
		}
	}
	return addr
}

// ConnectableString is a convenience wrapper that returns ConnectableAddr(addr).String().
func ConnectableString(addr net.Addr) string {
	return ConnectableAddr(addr).String()
}

var (
	// IPv4loopback is the loopback IP address for IPv4.
	IPv4loopback = net.IPv4(127, 0, 0, 1)
	// IPv6loopback is the loopback IP address for IPv6.
	IPv6loopback = net.IPv6loopback
)

// LoopbackIP returns the loopback IP address for the given IP address.
func LoopbackIP(ip net.IP) net.IP {
	if ip.To4() != nil {
		return IPv4loopback
	}
	return IPv6loopback
}
