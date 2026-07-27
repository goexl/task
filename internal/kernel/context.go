package kernel

import (
	"context"
)

type Context struct {
	context.Context
}

func NewContext(ctx context.Context) *Context {
	return &Context{
		Context: ctx,
	}
}

func (c *Context) Put(key any, value any) {
	c.Context = context.WithValue(c.Context, key, value)
}
