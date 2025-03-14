package core

import (
	"context"
	"sync"
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/gox/rand"
	"github.com/goexl/task/internal/internal/constant"
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
	defer func() {
		err = p.cleanup(ctx, task, &err)
	}()

	if re := p.updateRunning(task); nil != re {
		err = re
	} else if executor, pe := selector.Select(task); nil != pe {
		err = pe
	} else {
		err = executor.Execute(ctx, task.Target(), task.Times())
	}

	return
}

func (p *Processor) cleanup(ctx *kernel.Context, task kernel.Task, result *error) (err error) {
	// 从处理队列移除
	p.progresses.Delete(task.Id())

	if nil == *result { // 执行成功
		err = p.success(ctx, task)
	} else if task.Times() >= task.Maximum() {
		err = p.tasker.Failed(task)
	} else { // 执行失败
		// 确定下一次重试的时间，计算规则是，以二的幂为基数重试
		maximum := time.Second << task.Times()
		from := maximum * 4 / 5
		if maximum < p.minimize { // 如果不足最小间隔时间，在原来时间的基础上加最小间隔时间，确保随机出来的时间尽量分布均匀一点
			maximum = maximum + p.minimize
			from = maximum / 4
		}
		fixed := rand.New().Duration().Between(from, maximum).Build().Generate() // 随机重试，避免大量任务在同一时间重试
		runtime := time.Now().Add(fixed)
		err = p.tasker.Update(task.Id(), kernel.StatusFailed, runtime)
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

func (p *Processor) success(ctx *kernel.Context, task kernel.Task) (err error) {
	if kernel.TypeComputable == task.Type() { // 计算任务，将下一次执行时间交给处理器自身
		err = p.tasker.Update(task.Id(), kernel.StatusSuccess, ctx.Value(constant.KeyRuntime).(time.Time))
	} else if kernel.TypeFixed == task.Type() { // 存档
		err = p.tasker.Archive(task)
	} else {
		err = p.tasker.Update(task.Id(), kernel.StatusSuccess, task.Next())
	}

	return
}
