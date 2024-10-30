package core

import (
	"context"
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/task/internal/internal/constant"
	kernel2 "github.com/goexl/task/internal/internal/kernel"
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
			go func() {
				_ = p.process(_task, selector) // 错误已经处理，纯接收
			}()
		}
	}
}

func (p *Processor) process(task kernel.Task, selector kernel.Selector) (err error) {
	ctx := kernel2.NewContext(context.Background())
	defer func() {
		err = p.cleanup(ctx, task, &err)
	}()

	if re := p.updateRunning(task); nil != re {
		err = re
	} else if executor, pe := selector.Select(task); nil != pe {
		err = pe
	} else {
		err = executor.Execute(ctx, task.Target(), task.Retries())
	}

	return
}

func (p *Processor) cleanup(ctx *kernel2.Context, task kernel.Task, result *error) (err error) {
	if nil == *result { // 执行成功
		err = p.success(ctx, task)
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

func (p *Processor) success(ctx *kernel2.Context, task kernel.Task) (err error) {
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
