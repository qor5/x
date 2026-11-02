package errornotifierx

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"

	"github.com/pkg/errors"
	"github.com/theplant/inject/lifecycle"
)

func SetupLogNotifier(lc *lifecycle.Lifecycle) (Notifier, error) {
	notifier := NewLogNotifier()

	lc.Add(lifecycle.NewFuncActor(nil, func(_ context.Context) error {
		if err := notifier.Close(); err != nil {
			return errors.Wrap(err, "failed to close log notifier")
		}
		return nil
	}).WithName("log-notifier"))

	return notifier, nil
}

func NewLogNotifier() *LogNotifier {
	return &LogNotifier{
		logger: slog.Default(),
	}
}

type LogNotifier struct {
	logger *slog.Logger
}

func (n *LogNotifier) WithLogger(logger *slog.Logger) *LogNotifier {
	n.logger = logger
	return n
}

func (n *LogNotifier) Notify(val any, req *http.Request, logCtx map[string]any) {
	var ctx context.Context
	if req != nil {
		ctx = req.Context()
	} else {
		ctx = context.Background()
	}
	n.logger.ErrorContext(
		ctx,
		fmt.Sprintf("Error notification: %v", val),
		"err", val,
		"context", fmt.Sprint(logCtx),
		"stacktrace", string(debug.Stack()),
	)
}

func (n *LogNotifier) Close() error {
	return nil
}
