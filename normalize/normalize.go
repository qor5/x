package normalize

import (
	"context"
	"net"
	"strings"
	"sync"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"github.com/pkg/errors"
)

// Deprecated: use ClientKindUndefined instead
type ClientKind string

const (
	ClientKindPublic    ClientKind = "PUBLIC"
	ClientKindPrivate   ClientKind = "PRIVATE"
	ClientKindUndefined ClientKind = "UNDEFINED"
)

func (k ClientKind) IsPrivate() bool {
	return k == ClientKindPrivate
}

type CallMeta struct {
	ClientKind ClientKind
	Service    any
	FullMethod string
	Req        any
}

type normalization struct {
	cache     sync.Map
	callMeta  *CallMeta // for grpc or connect
	httpMeta  *HTTPMeta // for http
	setHeader func(key, value string)
}

type ctxKeyNormalization struct{}

func fromContext(ctx context.Context) (*normalization, error) {
	cm, ok := ctx.Value(ctxKeyNormalization{}).(*normalization)
	if !ok {
		return nil, errors.New("normalize: normalization not found in context, please configure the interceptor properly")
	}
	return cm, nil
}

func mustFromContext(ctx context.Context) *normalization {
	n, err := fromContext(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func CallMetaFromContext(ctx context.Context) (*CallMeta, error) {
	nlz, err := fromContext(ctx)
	if err != nil {
		return nil, err
	}
	return nlz.callMeta, nil
}

func MustCallMetaFromContext(ctx context.Context) *CallMeta {
	return mustFromContext(ctx).callMeta
}

func MustSetHeader(ctx context.Context, key, value string) {
	mustFromContext(ctx).setHeader(key, value)
}

func MustStore(ctx context.Context, key, value any) {
	mustFromContext(ctx).cache.Store(key, value)
}

func MustLoad(ctx context.Context, key any) (value any, ok bool) {
	return mustFromContext(ctx).cache.Load(key)
}

func ClientIPFromContext(ctx context.Context) string {
	md := metadata.ExtractIncoming(ctx)
	if md == nil {
		return ""
	}

	headerValues := [][]string{md["x-real-ip"], md["x-forwarded-for"]}
	for _, values := range headerValues {
		for _, value := range values {
			for _, addr := range strings.Split(value, ",") {
				addr = strings.TrimSpace(addr)
				if addr == "" {
					continue
				}
				ipaddr := net.ParseIP(addr)
				if ipaddr != nil {
					return ipaddr.String()
				}
				if clientIP, _, err := net.SplitHostPort(addr); err == nil {
					return clientIP
				}
			}
		}
	}
	return ""
}
