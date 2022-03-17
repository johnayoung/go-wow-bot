package ui

import (
	"fmt"
	"sort"

	"github.com/johnayoung/go-wow-bot/state"
	"github.com/jroimartin/gocui"
)

const (
	StateView = "State"
	BuffsView = "Buffs"
	SpellView = "Spells"
)

const NumGoroutines = 10

type UI struct {
	// g            *gocui.Gui
	Errors       chan<- error
	StateResults <-chan state.GameState
}

func (ui *UI) Start(g *gocui.Gui) {
	for state := range ui.StateResults {
		ui.update(g, state)
	}
}

func (ui *UI) update(g *gocui.Gui, state state.GameState) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View(StateView)
		if err != nil {
			return err
		}
		v.Clear()
		ui.show(v, "HealthMax", state.HealthMax)
		ui.show(v, "HealthCurrent", state.HealthCurrent)
		ui.show(v, "Health", state.Health)
		ui.show(v, "ManaMax", state.ManaMax)
		ui.show(v, "ManaCurrent", state.ManaCurrent)
		ui.show(v, "Mana", state.Mana)
		ui.show(v, "LowMana", state.LowMana)
		ui.show(v, "EnergyCurrent", state.EnergyCurrent)
		ui.show(v, "EnergyMax", state.EnergyMax)
		ui.show(v, "Energy", state.Energy)
		ui.show(v, "ComboPoints", state.ComboPoints)
		ui.show(v, "RageCurrent", state.RageCurrent)
		ui.show(v, "RageMax", state.RageMax)

		v, err = g.View(BuffsView)
		if err != nil {
			return err
		}
		v.Clear()
		ui.showBuffs(v, state.Buffs)
		ui.showDebuffs(v, state.Debuffs)
		ui.showDebuffs(v, state.TargetDebuffs)

		v, err = g.View(SpellView)
		if err != nil {
			return err
		}
		v.Clear()
		ui.showSpells(v, state.Spells)

		return nil
	})
}

func (ui *UI) show(v *gocui.View, name string, value interface{}) {
	fmt.Fprintln(v, name, value)
}

func (ui *UI) showBuffs(v *gocui.View, m map[string]state.Buff) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		ui.show(v, k, m[k])
	}
}

func (ui *UI) showDebuffs(v *gocui.View, m map[string]state.Debuff) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		ui.show(v, k, m[k])
	}
}

func (ui *UI) showSpells(v *gocui.View, m map[string]state.SpellMeta) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		ui.show(v, k, m[k])
	}
}
