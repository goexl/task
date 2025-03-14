package param

import (
	"github.com/goexl/log"
)

type Agent struct {
	Logger log.Logger `json:"logger,omitempty"`
}

func NewAgent() *Agent {
	return &Agent{
		Logger: log.New().Apply(),
	}
}
