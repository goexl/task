package builder

import (
	"github.com/goexl/log"
	"github.com/goexl/task/internal/core"
	"github.com/goexl/task/internal/param"
)

type Agent struct {
	params *param.Agent
}

func NewAgent() *Agent {
	return &Agent{
		params: param.NewAgent(),
	}
}

func (a *Agent) Logger(logger log.Logger) (agent *Agent) {
	a.params.Logger = logger
	agent = a

	return
}

func (a *Agent) Build() *core.Agent {
	return core.NewAgent(a.params)
}
