package param

import (
	"github.com/goexl/log"
	"github.com/goexl/task/internal/internal/kernel"
)

type Agent struct {
	Retries uint32        `json:"retries,omitempty"`
	Tasker  kernel.Tasker `json:"tasker,omitempty"`
	Logger  log.Logger    `json:"logger,omitempty"`
}

func NewAgent(tasker kernel.Tasker) *Agent {
	return &Agent{
		Retries: 10,
		Tasker:  tasker,
		Logger:  log.New().Apply(),
	}
}
