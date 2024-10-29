package core

import (
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/task/internal/internal/constant"
	"github.com/goexl/task/internal/kernel"
	"github.com/goexl/task/internal/param"
	"github.com/robfig/cron/v3"
)

type Processor struct {
	parser cron.Parser
	params *param.Agent
}

func NewProcessor(params *param.Agent) *Processor {
	return &Processor{
		params: params,
	}
}

func (p *Processor) Process(selector kernel.Selector) {
	for {
		tasking, exists := p.params.Tasker.Pop(p.params.Retries)
		if !exists { // 让出时间切片
			time.Sleep(0)
			continue
		}

		p.process(tasking, selector)
	}
}

func (p *Processor) process(tasking Task, selector kernel.Selector) {
	var err error
	defer func() {
		err = p.cleanup(tasking, &err)
	}()

	if re := p.updateRunning(tasking); nil != re {
		err = re
	} else if executor, pe := selector.Select(tasking); nil != pe {
		err = pe
	} else {
		err = executor.Execute(tasking.Target(), tasking.Retries())
	}

	return
}

func (p *Processor) cleanup(tasking Task, result *error) (err error) {
	if nil == *result { // 执行成功
		err = p.success(tasking)
	} else if tasking.Retries() >= p.params.Retries {
		err = p.maxRetry(tasking)
	} else { // 执行失败
		// 确定下一次重试的时间，计算规则是，以二的幂为基数重试
		runtime := time.Now().Add(15 * time.Second * 2 << tasking.Retries())
		err = p.params.Tasker.Update(tasking.Id(), kernel.StatusFailed, runtime)
	}

	return
}

func (p *Processor) updateRunning(tasking Task) (err error) {
	id := tasking.Id()
	status := gox.Ift[kernel.Status](0 == tasking.Retries(), kernel.StatusRunning, kernel.StatusRetrying)
	retries := tasking.Retries() + 1
	err = p.params.Tasker.Running(id, status, retries)

	return
}

func (p *Processor) success(task Task) (err error) {
	if p.removable(task) { // 是否删除任务本身
		err = p.params.Tasker.Remove(task)
	} else if kernel.TypeComputable == task.Type() { // 计算任务，将下一次执行时间交给处理器自身
		err = p.params.Tasker.Next(task.Id())
	} else {
		err = p.params.Tasker.Update(task.Id(), kernel.StatusSuccess, task.Next())
	}

	return
}

func (p *Processor) maxRetry(task Task) (err error) {
	if p.removable(task) {
		err = p.params.Tasker.Remove(task)
	}

	return
}

func (p *Processor) removable(task Task) (removable bool) {
	if cached, ok := task.Data()[constant.FieldRemovable]; ok {
		removable = cached.(bool)
	}

	return
}
