package core

import (
	"context"
	"sync"
	"time"

	"github.com/goexl/exception"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/gox/rand"
	"github.com/goexl/task/internal/kernel"
	"github.com/goexl/task/internal/param"
)

type Processor struct {
	tasker kernel.Tasker
	params *param.Agent

	progresses *sync.Map
	minimize   time.Duration // 任务执行最小间隔时间
}

func NewProcessor(tasker kernel.Tasker, params *param.Agent) *Processor {
	return &Processor{
		tasker: tasker,
		params: params,

		progresses: new(sync.Map),
		minimize:   5 * time.Second,
	}
}

func (p *Processor) Process(selector kernel.Selector) {
	for {
		_task := p.tasker.Pop()
		if _, exists := p.progresses.Load(_task.Id()); exists {
			time.Sleep(0) // 让出时间片
		} else {
			go func() {
				_ = p.process(_task, selector) // 错误已经处理，纯接收
			}()
		}
	}
}

func (p *Processor) process(task kernel.Task, selector kernel.Selector) (err error) {
	// 放入处理缓存
	p.progresses.Store(task.Id(), new(gox.Empty))

	ctx := kernel.NewContext(context.Background())
	var executor *kernel.Executor
	defer func() {
		err = p.cleanup(ctx, task, executor, &err)
	}()

	if re := p.updateRunning(task); nil != re {
		err = re
	} else if selected, pe := selector.Select(task); nil != pe {
		err = pe
	} else {
		executor = &selected
		err = selected.Execute(ctx, task)
	}

	return
}

func (p *Processor) cleanup(
	ctx *kernel.Context, task kernel.Task, executor *kernel.Executor, result *error,
) (err error) {
	// 从处理队列移除
	p.progresses.Delete(task.Id())

	if nil == *result { // 执行成功
		err = p.success(ctx, result, task, executor)
	} else if maximum := task.Maximum(); 0 != maximum && task.Times() >= maximum {
		err = p.tasker.Failed(task)
	} else { // 执行失败
		err = p.tasker.Update(task.Id(), kernel.StatusFailed, p.nextTime(ctx, result, task, executor))
	}

	return
}

func (p *Processor) updateRunning(task kernel.Task) (err error) {
	id := task.Id()
	status := gox.Ift[kernel.Status](0 == task.Times(), kernel.StatusRunning, kernel.StatusRetrying)
	retries := task.Times() + 1
	err = p.tasker.Running(id, status, retries)

	return
}

func (p *Processor) nextTime(
	ctx *kernel.Context, err *error, task kernel.Task, executor *kernel.Executor,
) (runtime time.Time) {
	switch {
	case kernel.TypeComputable == task.Type() && err == nil:
		runtime = *(*executor).(kernel.NextTimer).Next(ctx, nil, task)
	case kernel.TypeComputable == task.Type() && err != nil: // 计算任务，将下一次执行时间交给处理器自身
		runtime = *(*executor).(kernel.NextTimer).Next(ctx, *err, task)
	case executor == nil:
		runtime = time.Now().Add(time.Second)
	default:
		runtime = p.calNextTime(task)
	}

	return
}

func (p *Processor) calNextTime(task kernel.Task) (runtime time.Time) {
	// 确定下一次重试的时间，计算规则是，以二的幂为基数重试
	maximum := time.Second << task.Times()
	from := maximum * 4 / 5
	if maximum < p.minimize { // 如果不足最小间隔时间，在原来时间的基础上加最小间隔时间，确保随机出来的时间尽量分布均匀一点
		maximum = maximum + p.minimize
		from = maximum / 4
	}
	fixed := rand.New().Duration().Between(from, maximum).Build().Generate() // 随机重试，避免大量任务在同一时间重试
	runtime = time.Now().Add(fixed)

	return
}

func (p *Processor) success(ctx *kernel.Context, result *error, task kernel.Task, executor *kernel.Executor) (err error) {
	switch task.Type() {
	case kernel.TypeCron, kernel.TypeRate:
		err = p.tasker.Update(task.Id(), kernel.StatusStandby, task.Next())
	case kernel.TypeComputable:
		if next := (*executor).(kernel.NextTimer).Next(ctx, *result, task); next != nil {
			err = p.tasker.Update(task.Id(), kernel.StatusStandby, p.nextTime(ctx, result, task, executor))
		} else {
			err = p.tasker.Archive(task)
		}
	case kernel.TypeUnknown:
		err = exception.New().Message("没有匹配的类型").Field(field.New("type", "unknown")).Build()
	default:
		err = p.tasker.Archive(task)
	}

	return
}
