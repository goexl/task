package core

import (
	"context"
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/task"
	"github.com/goexl/task/internal/internal/constant"
	"github.com/goexl/task/internal/kernel"
	"github.com/goexl/task/internal/param"
)

type Processor struct {
	tasker kernel.Tasker
	params *param.Agent
}

func NewProcessor(tasker kernel.Tasker, params *param.Agent) *Processor {
	return &Processor{
		tasker: tasker,
		params: params,
	}
}

func (p *Processor) Process(selector kernel.Selector) {
	for {
		tasks, exists := p.tasker.Pop(p.params.Retries)
		if !exists { // 让出时间切片
			time.Sleep(0)
			continue
		}

		for _, _task := range tasks {
			go p.process(_task, selector)
		}
	}
}

func (p *Processor) process(_task kernel.Task, selector kernel.Selector) {
	ctx := NewContext(context.Background())
	var err error
	defer func() {
		err = p.cleanup(_task, &executor, &err)
	}()

	if re := p.updateRunning(_task); nil != re {
		err = re
	} else if selected, pe := selector.Select(_task); nil != pe {
		err = pe
	} else {
		executor = selected
		err = executor.Execute(ctx, _task.Target(), _task.Retries())
	}

	return
}

func (p *Processor) cleanup(ctx *Context, task kernel.Task, result *error) (err error) {
	if nil == *result { // 执行成功
		err = p.success(task, executor)
	} else if task.Retries() >= p.params.Retries {
		err = p.maxRetry(task)
	} else { // 执行失败
		// 确定下一次重试的时间，计算规则是，以二的幂为基数重试
		runtime := time.Now().Add(15 * time.Second * 2 << task.Retries())
		err = p.tasker.Update(task.Id(), kernel.StatusFailed, runtime)
	}

	return
}

func (p *Processor) updateRunning(tasking kernel.Task) (err error) {
	id := tasking.Id()
	status := gox.Ift[kernel.Status](0 == tasking.Retries(), kernel.StatusRunning, kernel.StatusRetrying)
	retries := tasking.Retries() + 1
	err = p.tasker.Running(id, status, retries)

	return
}

func (p *Processor) success(ctx *Context, task kernel.Task, executor *task.Executor) (err error) {
	if p.removable(task) { // 是否删除任务本身
		err = p.tasker.Remove(task)
	} else if kernel.TypeComputable == task.Type() { // 计算任务，将下一次执行时间交给处理器自身
		err = p.tasker.Update(task.Id(), kernel.StatusSuccess, ctx.Value(constant.KeyRuntime).(time.Time))
	} else {
		err = p.tasker.Update(task.Id(), kernel.StatusSuccess, task.Next())
	}

	return
}

func (p *Processor) maxRetry(task kernel.Task) (err error) {
	if p.removable(task) {
		err = p.tasker.Remove(task)
	}

	return
}

func (p *Processor) removable(task kernel.Task) (removable bool) {
	if cached, ok := task.Data()[constant.FieldRemovable]; ok {
		removable = cached.(bool)
	}

	return
}
