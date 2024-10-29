package param

import (
	"github.com/goexl/log"
)

type Agent struct {
	Retries uint32     `json:"retries,omitempty"`
	Logger  log.Logger `json:"logger,omitempty"`
}

func NewAgent() *Agent {
	return &Agent{
		Retries: 10,
		Logger:  log.New().Apply(),
	}
}
