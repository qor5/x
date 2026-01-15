package netx

import (
	"fmt"
	"net"
)

// ConnectableString converts a net.Addr to a string address that can be used for connecting.
// It handles various edge cases:
//   - Unspecified addresses (0.0.0.0, [::]) are converted to loopback (127.0.0.1, [::1])
//   - IPv6 addresses are properly formatted with brackets for host:port usage
//   - Unix socket addresses are returned as-is
//   - Preserves the original address family (IPv4 vs IPv6)
func ConnectableString(addr net.Addr) string {
	switch a := addr.(type) {
	case *net.TCPAddr:
		return connectableTCPAddrString(a)
	case *net.UDPAddr:
		return connectableUDPAddrString(a)
	case *net.UnixAddr:
		return a.String()
	case *net.IPAddr:
		return connectableIPAddrString(a)
	default:
		return addr.String()
	}
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

func connectableTCPAddrString(addr *net.TCPAddr) string {
	ip := addr.IP
	if ip.IsUnspecified() {
		ip = LoopbackIP(ip)
	}
	return formatIPPort(ip, addr.Port, addr.Zone)
}

func connectableUDPAddrString(addr *net.UDPAddr) string {
	ip := addr.IP
	if ip.IsUnspecified() {
		ip = LoopbackIP(ip)
	}
	return formatIPPort(ip, addr.Port, addr.Zone)
}

func connectableIPAddrString(addr *net.IPAddr) string {
	ip := addr.IP
	if ip.IsUnspecified() {
		ip = LoopbackIP(ip)
	}
	return formatIP(ip, addr.Zone)
}

func formatIPPort(ip net.IP, port int, zone string) string {
	if ip.To4() != nil {
		return fmt.Sprintf("%s:%d", ip.String(), port)
	}

	ipStr := ip.String()
	if zone != "" {
		ipStr = fmt.Sprintf("%s%%%s", ipStr, zone)
	}
	return fmt.Sprintf("[%s]:%d", ipStr, port)
}

func formatIP(ip net.IP, zone string) string {
	if ip.To4() != nil {
		return ip.String()
	}

	ipStr := ip.String()
	if zone != "" {
		return fmt.Sprintf("%s%%%s", ipStr, zone)
	}
	return ipStr
}
