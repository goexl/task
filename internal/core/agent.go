package core

import (
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/log"
	"github.com/goexl/task"
	"github.com/goexl/task/internal/internal/constant"
	"github.com/goexl/task/internal/internal/core"
	"github.com/goexl/task/internal/kernel"
	"github.com/goexl/task/internal/param"
	"github.com/robfig/cron/v3"
)

type Agent struct {
	parser cron.Parser
	params *param.Task
	tasker core.Tasker
	logger log.Logger
}

func NewAgent(params *param.Task, tasker core.Tasker, logger log.Logger) *Agent {
	return &Agent{
		parser: cron.NewParser(cron.Second),
		params: params,
		tasker: tasker,
		logger: logger,
	}
}

func (a *Agent) Process(tasking core.Tasking, processor core.Processor) (err error) {
	defer func() {
		err = a.cleanup(tasking, &err)
	}()

	if re := a.updateRunning(tasking); nil != re {
		err = re
	} else if executor, pe := processor.Process(tasking); nil != pe {
		err = pe
	} else {
		err = executor.Execute(tasking.Target(), tasking.Retries())
	}

	return
}

func (a *Agent) Add(scheduling core.Scheduling) error {
	return a.tasker.Add(scheduling)
}

func (a *Agent) Remove(scheduling core.Scheduling) error {
	return a.tasker.Remove(scheduling)
}

func (a *Agent) run(tasking core.Tasking) (err error) {

	return
}

func (a *Agent) updateRunning(tasking core.Tasking) (err error) {
	id := tasking.Id()
	status := gox.Ift[task.Status](0 == tasking.Retries(), task.StatusRunning, task.StatusRetrying)
	retries := tasking.Retries() + 1
	err = a.tasker.Running(id, status, retries)

	return
}

func (a *Agent) cleanup(tasking core.Tasking, result *error) (err error) {
	if nil == *result { // 执行成功
		err = a.success(tasking)
	} else if tasking.Retries() >= a.params.Retries {
		err = a.maxRetry(tasking)
	} else { // 执行失败
		// 确定下一次重试的时间，计算规则是，以二的幂为基数重试
		runtime := time.Now().Add(15 * time.Second * 2 << tasking.Retries())
		err = a.tasker.Update(tasking.Id(), task.StatusFailed, runtime)
	}

	return
}

func (a *Agent) success(tasking core.Tasking) (err error) {
	if tasking.Removable() { // 是否删除任务本身
		err = a.tasker.Remove(tasking)
	} else if kernel.TypeComputable == tasking.Type() { // 计算任务，将下一次执行时间交给处理器自身
		err = a.tasker.Next(tasking.Id())
	} else if kernel.TypeOnce == tasking.Type() { // 只执行一次的任务，在执行成功后，将时间调整到5小时前，避免重复执行
		err = a.tasker.Update(tasking.Id(), task.StatusSuccess, time.Now().Add(-5*time.Hour))
	} else if kernel.TypeRate == tasking.Type() { // 周期性任务，按要求计算下一次执行时间
		runtime := time.Now().Add(tasking.Data()[constant.FieldRate].(time.Duration))
		err = a.tasker.Update(tasking.Id(), task.StatusSuccess, runtime)
	} else if kernel.TypeCron == tasking.Type() { // 计划任务
		runtime := a.nextCron(tasking)
		err = a.tasker.Update(tasking.Id(), task.StatusSuccess, runtime)
	}

	return
}

func (a *Agent) maxRetry(tasking core.Tasking) (err error) {
	if tasking.Removable() {
		err = a.tasker.Remove(tasking)
	}

	return
}

func (a *Agent) nextCron(tasking core.Tasking) (next time.Time) {
	expression := tasking.Data()[constant.FieldCron].(string)
	if schedule, pe := a.parser.Parse(expression); nil != pe {
		next = time.Now()
	} else {
		next = schedule.Next(time.Now())
	}

	return
}
