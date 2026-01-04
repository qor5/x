package httptestx

import (
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/theplant/inject/lifecycle"
)

func SetupServer(lc *lifecycle.Lifecycle, handler http.Handler) (*httptest.Server, error) {
	server := httptest.NewServer(handler)
	lc.Add(lifecycle.NewFuncActor(nil, func(_ context.Context) error {
		server.Close()
		return nil
	}).WithName("httptest-server"))
	return server, nil
}
