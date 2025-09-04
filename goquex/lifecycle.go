package goquex

import (
	"context"

	"github.com/pkg/errors"
	"github.com/qor5/go-bus/quex"
)

type WorkerService struct {
	quex.WorkerController
	startFunc func(ctx context.Context) (quex.WorkerController, error)
	name      string
}

func NewWorkerService(startFunc func(ctx context.Context) (quex.WorkerController, error)) *WorkerService {
	return &WorkerService{
		startFunc: startFunc,
	}
}

func (q *WorkerService) Start(ctx context.Context) error {
	ctrl, err := q.startFunc(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to start worker controller")
	}
	q.WorkerController = ctrl
	return nil
}

func (q *WorkerService) WithName(name string) *WorkerService {
	q.name = name
	return q
}

func (q *WorkerService) GetName() string {
	return q.name
}
