package core

import (
	"context"

	"github.com/goexl/task/internal/internal/core"
	"github.com/goexl/task/internal/kernel"
	"github.com/goexl/task/internal/param"
)

type Agent struct {
	params *param.Agent
	tasker kernel.Tasker
}

func NewAgent(params *param.Agent) *Agent {
	return &Agent{
		params: params,
	}
}

func (a *Agent) Start(ctx context.Context, tasker kernel.Tasker, selector kernel.Selector) (err error) {
	if err = tasker.Start(ctx); nil == err {
		a.tasker = tasker

		processor := core.NewProcessor(tasker, a.params)
		go processor.Process(selector)
	}

	return
}

func (a *Agent) Add(schedule kernel.Schedule) error {
	return a.tasker.Add(schedule)
}

func (a *Agent) Remove(schedule kernel.Schedule) error {
	return a.tasker.Remove(schedule)
}

func (a *Agent) Stop(ctx context.Context) error {
	return a.tasker.Stop(ctx)
}
