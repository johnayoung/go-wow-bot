package run

import (
	"log"
	"time"

	"github.com/johnayoung/go-wow-bot/state"
)

type Agent struct {
	Status       string
	UpdateTicker <-chan time.Time
	Errors       chan<- error
	StateManager state.StateManager
	StateResults chan<- state.GameState
}

func (a *Agent) Start() {
	log.Print("Starting agent.")

	for {
		select {
		case <-a.UpdateTicker:
			err := a.evaluate()
			if err != nil {
				a.Errors <- err
			}
		}
	}
}

func (a *Agent) evaluate() error {
	state := a.StateManager.Get()

	a.StateResults <- state
	return nil
}
