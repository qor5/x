package goquex

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"github.com/qor5/go-que"
	"github.com/qor5/x/v3/jsonx"
	"github.com/qor5/x/v3/panicx"
	"github.com/theplant/appkit/errornotifier"
	"github.com/theplant/appkit/logtracing"
)

var PlansSampleSize = 10

type withTracing struct {
	que.Queue
}

func (g *withTracing) Enqueue(ctx context.Context, tx *sql.Tx, plans ...que.Plan) (_ []int64, xerr error) {
	spanName := "goque.enqueue"
	if len(plans) == 1 {
		spanName = fmt.Sprintf("goque.enqueue:%s", plans[0].Queue)
	}
	ctx, span := logtracing.StartSpan(ctx, spanName)
	defer func() { logtracing.EndSpan(ctx, xerr) }()

	span.AppendKVs("goque.plans.count", len(plans))
	span.AppendKVs("goque.transaction.enabled", tx != nil)

	if len(plans) > 0 {
		span.AppendKVs("goque.first_queue.name", plans[0].Queue)

		var payload []byte
		var err error

		if len(plans) > PlansSampleSize {
			span.AppendKVs("goque.plans.oversized", true)
			span.AppendKVs("goque.plans.sample_size", PlansSampleSize)
			payload, err = jsonx.Marshal(plans[:PlansSampleSize])
		} else {
			payload, err = jsonx.Marshal(plans)
		}

		if err != nil {
			span.AppendKVs("goque.plans.serialize_error", err.Error())
			span.AppendKVs("goque.plans.serialize_failed", true)
		} else {
			span.AppendKVs("goque.plans.data", string(payload))
			span.AppendKVs("goque.plans.data_size", len(payload))
		}
	}

	ids, err := g.Queue.Enqueue(ctx, tx, plans...)
	if err != nil {
		return nil, errors.Wrap(err, "goque.enqueue failed")
	}
	return ids, nil
}

func WithTracing(goq que.Queue) que.Queue {
	return &withTracing{
		Queue: goq,
	}
}

func PerformWithTracing(notifier errornotifier.Notifier) func(next func(ctx context.Context, j que.Job) error) func(ctx context.Context, j que.Job) error {
	return func(next func(ctx context.Context, j que.Job) error) func(ctx context.Context, j que.Job) error {
		return func(ctx context.Context, j que.Job) (xerr error) {
			spanName := fmt.Sprintf("goque.perform:%s", j.Plan().Queue)
			ctx, span := logtracing.StartSpan(ctx, spanName)
			defer func() { logtracing.EndSpan(ctx, xerr) }()
			defer logtracing.RecordPanic(ctx)

			spanKVs := map[string]any{
				"goque.job.id":     j.ID(),
				"goque.queue.name": j.Plan().Queue,
			}

			payload, err := jsonx.Marshal(j.Plan())
			if err != nil {
				spanKVs["goque.plan.serialize_error"] = err.Error()
				spanKVs["goque.plan.serialize_failed"] = true
			} else {
				spanKVs["goque.plan.data"] = string(payload)
				spanKVs["goque.plan.data_size"] = len(payload)
			}

			for k, v := range spanKVs {
				span.AppendKVs(k, v)
			}

			if notifier != nil {
				defer panicx.On(func(perr error) {
					if perr == nil {
						perr = xerr
					}
					if perr == nil {
						return
					}
					// add more context to the error notification
					errorContext := make(map[string]any)
					for k, v := range spanKVs {
						errorContext[k] = v
					}
					errorContext["goque.span.name"] = spanName
					errorContext["goque.error.source"] = "job_execution"
					notifier.Notify(perr, nil, errorContext)
				})
			}

			return next(ctx, j)
		}
	}
}
