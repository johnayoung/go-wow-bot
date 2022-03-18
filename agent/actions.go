package agent

import (
	"github.com/go-vgo/robotgo"
	"github.com/johnayoung/go-wow-bot/state"
)

var Actions = []Action{
	// Combo point
	CreateComboPointAction("sinisterStrike", 9, 45, 4, false, nil),
	CreateComboPointAction("hemorrhage", 9, 27, 4, false, nil),
	CreateComboPointAction("eviscerate", 9, 27, 4, true, nil),

	// Rage
	CreateRageAction("slam", 8, 11, nil),
	CreateRageAction("mortalStrike", 7, 22, nil),
	CreateRageAction("dragonRoar", 5, 1, nil),
	CreateRageAction("execute", 7, 11, func(s state.GameState) bool {
		if playerHasBuff("suddenDeath", s) {
			return inMeleeRange(s)
		}
		return false
	}),

	// Mana
	CreateManaAction("serpentSting", 7, 10, func(s state.GameState) bool {
		return !inMeleeRange(s) && !targetHasDebuff("serpentSting", s)
	}),
	CreateManaAction("locustShot", 8, 68, func(s state.GameState) bool {
		return !inMeleeRange(s)
	}),
	CreateManaAction("wrath", 10, 136, func(s state.GameState) bool {
		return playerHasBuff("wildWrath", s)
	}),
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
	return conditionsMet && !effectsAchieved
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

func CreateManaAction(
	name string,
	cost,
	manaCost int,
	conditions func(s state.GameState) bool,
) *DefaultAction {
	return CreateAction(
		name,
		func(s state.GameState) int {
			if name == "wrath" && playerHasBuff("wildWrath", s) {
				return 6
			}
			return cost
		},
		func(s state.GameState) bool {
			spellAvailable := available(name, s)
			haveMana := int(s.ManaCurrent) > manaCost
			extra := true

			if conditions != nil {
				extra = addConditions(s, conditions)
			}

			return s.MemberCombatStatus["playerInCombat"] &&
				spellAvailable &&
				haveMana &&
				extra
		},
		targetIsDead,
	)
}

func CreateRageAction(
	name string,
	cost,
	rageCost int,
	conditions func(s state.GameState) bool,
) *DefaultAction {
	return CreateAction(
		name,
		func(s state.GameState) int {
			if name == "execute" && playerHasBuff("suddenDeath", s) {
				return 6
			}
			return cost
		},
		func(s state.GameState) bool {
			spellAvailable := available(name, s)
			haveRage := int(s.RageCurrent) > rageCost
			extra := true

			if conditions != nil {
				extra = addConditions(s, conditions)
			}

			return spellAvailable &&
				inMeleeRange(s) &&
				haveRage &&
				extra
		},
		targetIsDead,
	)
}

func CreateComboPointAction(
	name string,
	cost,
	energyCost,
	minComboPoints int,
	spend bool,
	conditions func(s state.GameState) bool,
) *DefaultAction {
	return CreateAction(
		name,
		func(s state.GameState) int {
			return cost
		},
		func(s state.GameState) bool {
			cpLogic := s.ComboPoints < int64(minComboPoints)
			if spend {
				cpLogic = s.ComboPoints >= int64(minComboPoints)
			}

			extra := true

			if conditions != nil {
				extra = addConditions(s, conditions)
			}

			return available(name, s) &&
				inMeleeRange(s) &&
				s.EnergyCurrent > int64(energyCost) &&
				cpLogic &&
				extra
		},
		targetIsDead,
	)
}

func targetIsDead(s state.GameState) bool {
	return s.MemberStatus["targetIsDead"]
}

func available(name string, s state.GameState) bool {
	spell := s.GetSpell(name)
	return spell.Castable && spell.Equipped
}

func addConditions(s state.GameState, conditions ...func(s state.GameState) bool) bool {
	for _, c := range conditions {
		valid := c(s)
		if valid == false {
			return false
		}
	}
	return true
}

func inMeleeRange(s state.GameState) bool {
	return s.MemberMeleeRange["targetInMeleeRange"] == true
}

func playerHasBuff(name string, s state.GameState) bool {
	hasBuff := false
	if _, ok := s.Buffs[name]; ok {
		//do something here
		hasBuff = true
	}
	return hasBuff
}

func targetHasDebuff(name string, s state.GameState) bool {
	hasDebuff := false
	if _, ok := s.TargetDebuffs[name]; ok {
		//do something here
		hasDebuff = true
	}
	return hasDebuff
}
