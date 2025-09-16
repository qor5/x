package grpcx

import (
	"net"
	"net/http"
	"strings"

	"google.golang.org/grpc/metadata"
)

var xForwardedFor = http.CanonicalHeaderKey("X-Forwarded-For")

func MetadataFromHeader(header http.Header, remoteAddr string) metadata.MD {
	var pairs []string
	for key, vals := range header {
		key = http.CanonicalHeaderKey(key)
		// nolint:gocritic
		switch key {
		case xForwardedFor:
			// Handled separately below
			continue
		}
		for _, val := range vals {
			pairs = append(pairs, key, val)
		}
	}

	xff := header.Values(xForwardedFor)
	if remoteAddr != "" {
		if ipaddr, _, err := net.SplitHostPort(remoteAddr); err == nil {
			xff = append(xff, ipaddr)
		}
	}
	if len(xff) > 0 {
		pairs = append(pairs, strings.ToLower(xForwardedFor), strings.Join(xff, ", "))
	}

	return metadata.Pairs(pairs...)
}
