package core

import (
	"context"

	"github.com/goexl/task/internal/internal/core"
	"github.com/goexl/task/internal/kernel"
	"github.com/goexl/task/internal/param"
)

type Agent struct {
	params    *param.Agent
	processor *core.Processor
}

func NewAgent(params *param.Agent) *Agent {
	return &Agent{
		params: params,
	}
}

func (a *Agent) Start(ctx context.Context, selector kernel.Selector) (err error) {
	if err = a.params.Tasker.Start(ctx); nil == err {
		go a.processor.Process(selector)
	}

	return
}

func (a *Agent) Add(scheduling core.Schedule) error {
	return a.params.Tasker.Add(scheduling)
}

func (a *Agent) Remove(scheduling core.Schedule) error {
	return a.params.Tasker.Remove(scheduling)
}

func (a *Agent) Stop(ctx context.Context) error {
	return a.params.Tasker.Stop(ctx)
}
