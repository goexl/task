package param

import (
	"github.com/goexl/log"
	"github.com/goexl/task/internal/kernel"
)

type Agent struct {
	Retries uint32     `json:"retries,omitempty"`
	Logger  log.Logger `json:"logger,omitempty"`

	Tasker kernel.Tasker `json:"tasker,omitempty"`
}

func NewAgent(tasker kernel.Tasker) *Agent {
	return &Agent{
		Retries: 10,
		Logger:  log.New().Apply(),

		Tasker: tasker,
	}
}
