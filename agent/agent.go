package agent

import (
	"fmt"
	"log"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/johnayoung/go-wow-bot/state"
)

type Agent struct {
	Status       chan<- string
	UpdateTicker <-chan time.Time
	Errors       chan<- error
	StateManager state.StateManager
	StateResults chan<- state.GameState
	Actions      []Action
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

	go func() {
		a.StateResults <- state
	}()

	// goal, _ := a.determineGoal(state)

	title := robotgo.GetTitle()
	if title != "Ascension" {
		return nil
	}

	// TODO: Perform
	go func() {
		action, _, err := a.GetBestAction(state)

		if err != nil {
			a.Status <- "No valid action found."
		} else {
			a.Status <- fmt.Sprintf("Casting %s", action.String())
			action.Run(state.FindKeybind(action.String()))
		}

	}()

	return nil
}

func (a *Agent) determineGoal(s state.GameState) (Goal, error) {
	if pull.Condition(s) {
		return pull, nil
	}

	if targetDead.Condition(s) {
		return targetDead, nil
	}

	return nil, fmt.Errorf("Unable to determine goal")
}

type Plan []Action

func (p *Plan) Cost(s state.GameState) int {
	cost := 0
	for _, action := range *p {
		cost += action.Cost(s)
	}
	return cost
}

func (a *Agent) possibleActions(s state.GameState) []Action {
	validActions := []Action{}
	for _, action := range a.Actions {
		if action.CanRun(s) {
			validActions = append(validActions, action)
		}
	}
	return validActions
}

func (a *Agent) GetBestAction(s state.GameState) (Action, int, error) {
	plans := a.possibleActions(s)

	if len(plans) == 0 {
		return nil, 0, fmt.Errorf("No possible actions found.")
	}

	var bestPlan Action = (*DefaultAction)(nil)
	bestCost := 99999
	for _, plan := range plans {
		cost := plan.Cost(s)
		if cost < bestCost {
			bestPlan = plan
			bestCost = cost
		}
	}
	return bestPlan, bestCost, nil
}
