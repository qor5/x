package ratelimiterx

import (
	"context"
	"time"

	"github.com/qor5/x/v3/jsonx"
	"github.com/qor5/x/v3/normalize"
	"github.com/qor5/x/v3/statusx"
	statusv1 "github.com/qor5/x/v3/statusx/gen/status/v1"
	"github.com/theplant/ratelimiter"
	"google.golang.org/grpc/codes"
)

var ErrorReasonRateLimited = "RATE_LIMITED"

type Evaluator func(ctx context.Context, callMeta *normalize.CallMeta) ([]*ratelimiter.ReserveRequest, error)

type Metadata struct {
	TimeToAct  time.Time `json:"timeToAct"`
	ReservedAt time.Time `json:"reservedAt"`
}

func allow(ctx context.Context, limiter ratelimiter.RateLimiter, evaluator Evaluator) error {
	callMeta := normalize.MustCallMetaFromContext(ctx)
	reserveRequests, err := evaluator(ctx, callMeta)
	if err != nil {
		return statusx.Wrap(err, codes.Internal, statusv1.ErrorReason_INTERNAL.String(), "failed to call policy function").Err()
	}

	if len(reserveRequests) > 0 {
		var meta Metadata
		for _, r := range reserveRequests {
			if r.MaxFutureReserve > 0 {
				return statusx.New(codes.Internal, statusv1.ErrorReason_INTERNAL.String(), "expect max future reserve is equal to 0").Err()
			}
			reservation, err := limiter.Reserve(ctx, r)
			if err != nil {
				return statusx.Wrapf(err, codes.Internal, statusv1.ErrorReason_INTERNAL.String(), "failed to reserve for key %s", r.Key).Err()
			}
			if reservation.OK {
				continue
			}
			if reservation.TimeToAct.After(meta.TimeToAct) {
				meta.TimeToAct = reservation.TimeToAct
				meta.ReservedAt = reservation.ReservedAt
			}
		}
		if !meta.TimeToAct.IsZero() {
			md := make(map[string]string)
			if err := jsonx.Copy(&md, &meta); err != nil {
				return statusx.Wrap(err, codes.Internal, statusv1.ErrorReason_INTERNAL.String(), "failed to copy metadata").Err()
			}
			return statusx.New(codes.ResourceExhausted, ErrorReasonRateLimited, "ratelimit exceeded").
				WithMetadata(md).
				Err()
		}
	}

	return nil
}
