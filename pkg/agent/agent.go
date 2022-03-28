package agent

import (
	"fmt"
	"log"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/johnayoung/go-wow-bot/state"
	hook "github.com/robotn/gohook"
)

type Agent struct {
	StateManager state.StateManager
	Paused       bool
	Status       chan<- string
	UpdateTicker <-chan time.Time
	StateResults chan<- state.GameState
	Errors       chan<- error
	Actions      []Action
}

func (a *Agent) Start() {
	log.Print("Starting agent.")

	lshift := uint16(160)

	hook.Register(hook.KeyHold, []string{}, func(e hook.Event) {
		isLeftShift := e.Rawcode == lshift

		if e.Kind == hook.KeyHold && isLeftShift {
			fmt.Println("Status: Paused")
		}
	})

	hook.Register(hook.KeyUp, []string{}, func(e hook.Event) {
		isLeftShift := e.Rawcode == lshift

		if e.Kind == hook.KeyUp && isLeftShift {
			fmt.Println("Status: Running")
		}
	})

	s := hook.Start()
	defer hook.End()

	for {
		select {
		case e := <-s:
			a.pause(e)
		case <-a.UpdateTicker:
			err := a.evaluate()
			if err != nil {
				a.Errors <- err
			}
		}
	}
}

func (a *Agent) pause(e hook.Event) {
	lshift := uint16(160)
	isLeftShift := e.Rawcode == lshift

	if e.Kind == hook.KeyHold && isLeftShift {
		a.Paused = true
	} else if e.Kind == hook.KeyUp && isLeftShift {
		a.Paused = false
	} else {
		// Do nothing
	}
}

func (a *Agent) evaluate() error {
	if a.Paused == true {
		return nil
	}

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

func (a *Agent) isLeftShiftHeldDown() error {
	return nil
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
