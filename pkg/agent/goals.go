package agent

import "github.com/johnayoung/go-wow-bot/state"

var (
	pull       = Pull{"Pull"}
	targetDead = TargetDead{"Target is dead"}
)

type Goal interface {
	Validate(state.GameState, state.GameState) bool
	Condition(state.GameState) bool
}

type Pull struct {
	Label string
}

func (p Pull) Validate(prevState state.GameState, nextState state.GameState) bool {
	return prevState.MemberCombatStatus["playerInCombat"] == false
}

func (p Pull) Condition(s state.GameState) bool {
	return s.MemberCombatStatus["playerInCombat"] == false
}

type TargetDead struct {
	Label string
}

func (t TargetDead) Validate(prevState state.GameState, nextState state.GameState) bool {
	return nextState.MemberStatus["targetIsDead"] == true
}

func (t TargetDead) Condition(s state.GameState) bool {
	return s.MemberCombatStatus["playerInCombat"] == true && s.MemberStatus["targetIsDead"] == false
}
