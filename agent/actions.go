package agent

import (
	"github.com/go-vgo/robotgo"
	"github.com/johnayoung/go-wow-bot/state"
)

var sinisterStrike = CreateAction(
	"sinisterStrike",
	func(s state.GameState) int {
		return 9
	},
	func(s state.GameState) bool {
		return s.EnergyCurrent > 45 && s.ComboPoints < 4
	},
	func(s state.GameState) bool {
		return s.MemberStatus["targetIsDead"]
	},
)

var actions = map[string]Action{
	"sinisterStrike": sinisterStrike,
}

type Action interface {
	String() string
	Cost(state.GameState) int
	CanRun(state.GameState) bool
	Run(string) string
}

// ensure DefaultAction implements Action
var _ Action = (*DefaultAction)(nil)

type Cost func(s state.GameState) int
type Assess func(s state.GameState) bool
type DefaultAction struct {
	name      string
	cost      Cost
	condition Assess
	effect    Assess
}

func (a *DefaultAction) String() string {
	return a.name
}

func (a *DefaultAction) Cost(s state.GameState) int {
	return a.cost(s)
}

func (a *DefaultAction) CanRun(s state.GameState) bool {
	conditionsMet := a.condition(s)
	effectsAchieved := a.effect(s)
	return conditionsMet && effectsAchieved
}

func (a *DefaultAction) Run(k string) string {
	return robotgo.KeyTap(k)
}

func CreateAction(
	name string,
	cost Cost,
	condition Assess,
	effect Assess,
) *DefaultAction {
	return &DefaultAction{
		name,
		cost,
		condition,
		effect,
	}
}
