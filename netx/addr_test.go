package netx

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectableString(t *testing.T) {
	tests := []struct {
		name     string
		addr     net.Addr
		expected string
	}{
		{
			name:     "IPv4 specific address",
			addr:     &net.TCPAddr{IP: net.ParseIP("192.168.1.1"), Port: 8080},
			expected: "192.168.1.1:8080",
		},
		{
			name:     "IPv4 unspecified address",
			addr:     &net.TCPAddr{IP: net.ParseIP("0.0.0.0"), Port: 8080},
			expected: "127.0.0.1:8080",
		},
		{
			name:     "IPv6 specific address",
			addr:     &net.TCPAddr{IP: net.ParseIP("2001:db8::1"), Port: 8080},
			expected: "[2001:db8::1]:8080",
		},
		{
			name:     "IPv6 unspecified address",
			addr:     &net.TCPAddr{IP: net.ParseIP("::"), Port: 8080},
			expected: "[::1]:8080",
		},
		{
			name:     "IPv6 with zone",
			addr:     &net.TCPAddr{IP: net.ParseIP("fe80::1"), Port: 8080, Zone: "eth0"},
			expected: "[fe80::1%eth0]:8080",
		},
		{
			name:     "IPv6 loopback",
			addr:     &net.TCPAddr{IP: net.ParseIP("::1"), Port: 8080},
			expected: "[::1]:8080",
		},
		{
			name:     "IPv4 loopback",
			addr:     &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8080},
			expected: "127.0.0.1:8080",
		},
		{
			name:     "UDP IPv4",
			addr:     &net.UDPAddr{IP: net.ParseIP("192.168.1.1"), Port: 53},
			expected: "192.168.1.1:53",
		},
		{
			name:     "UDP IPv4 unspecified",
			addr:     &net.UDPAddr{IP: net.ParseIP("0.0.0.0"), Port: 53},
			expected: "127.0.0.1:53",
		},
		{
			name:     "UDP IPv6",
			addr:     &net.UDPAddr{IP: net.ParseIP("2001:db8::1"), Port: 53},
			expected: "[2001:db8::1]:53",
		},
		{
			name:     "UDP IPv6 unspecified",
			addr:     &net.UDPAddr{IP: net.ParseIP("::"), Port: 53},
			expected: "[::1]:53",
		},
		{
			name:     "IPAddr IPv4",
			addr:     &net.IPAddr{IP: net.ParseIP("192.168.1.1")},
			expected: "192.168.1.1",
		},
		{
			name:     "IPAddr IPv4 unspecified",
			addr:     &net.IPAddr{IP: net.ParseIP("0.0.0.0")},
			expected: "127.0.0.1",
		},
		{
			name:     "IPAddr IPv6",
			addr:     &net.IPAddr{IP: net.ParseIP("2001:db8::1")},
			expected: "2001:db8::1",
		},
		{
			name:     "IPAddr IPv6 unspecified",
			addr:     &net.IPAddr{IP: net.ParseIP("::")},
			expected: "::1",
		},
		{
			name:     "IPAddr IPv6 with zone",
			addr:     &net.IPAddr{IP: net.ParseIP("fe80::1"), Zone: "eth0"},
			expected: "fe80::1%eth0",
		},
		{
			name:     "Unix socket",
			addr:     &net.UnixAddr{Name: "/tmp/test.sock", Net: "unix"},
			expected: "/tmp/test.sock",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConnectableString(tt.addr)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestLoopbackIP(t *testing.T) {
	tests := []struct {
		name     string
		ip       net.IP
		expected net.IP
	}{
		{
			name:     "IPv4 unspecified",
			ip:       net.ParseIP("0.0.0.0"),
			expected: IPv4loopback,
		},
		{
			name:     "IPv6 unspecified",
			ip:       net.ParseIP("::"),
			expected: IPv6loopback,
		},
		{
			name:     "IPv4 specific address",
			ip:       net.ParseIP("192.168.1.1"),
			expected: IPv4loopback,
		},
		{
			name:     "IPv6 specific address",
			ip:       net.ParseIP("2001:db8::1"),
			expected: IPv6loopback,
		},
		{
			name:     "IPv4 loopback",
			ip:       net.ParseIP("127.0.0.1"),
			expected: IPv4loopback,
		},
		{
			name:     "IPv6 loopback",
			ip:       net.ParseIP("::1"),
			expected: IPv6loopback,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LoopbackIP(tt.ip)
			assert.True(t, result.Equal(tt.expected))
		})
	}
}
