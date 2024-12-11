package kernel

import (
	"context"
	"time"

	"github.com/goexl/task/internal/internal/constant"
)

type Context struct {
	context.Context
}

func NewContext(ctx context.Context) *Context {
	return &Context{
		Context: ctx,
	}
}

func (c *Context) Runtime(runtime time.Time) {
	c.Context = context.WithValue(c.Context, constant.KeyRuntime, runtime)
}

func (c *Context) Put(key any, value any) {
	c.Context = context.WithValue(c.Context, key, value)
}
