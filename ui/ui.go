package ui

import (
	"fmt"
	"sort"

	"github.com/johnayoung/go-wow-bot/state"
	"github.com/jroimartin/gocui"
)

const (
	StateView         = "State"
	BuffsView         = "Buffs"
	DebuffsView       = "Debuffs"
	TargetDebuffsView = "TargetDebuffs"
	SpellView         = "Spells"
	StatusView        = "Status"
)

var Views = []string{
	StateView,
	BuffsView,
	DebuffsView,
	TargetDebuffsView,
	SpellView,
	StatusView,
}

const NumGoroutines = 10

type UI struct {
	Errors       chan<- error
	StateResults <-chan state.GameState
	Status       <-chan string
}

func (ui *UI) Start(g *gocui.Gui) {
	for {
		select {
		case state := <-ui.StateResults:
			ui.update(g, state)
		case status := <-ui.Status:
			// fmt.Sprintf("Received a status update: %s", status)
			ui.updateStatus(g, status)
		}
	}

	// for state := range ui.StateResults {
	// 	ui.update(g, state)
	// }

	// for status := range ui.Status {
	// 	fmt.Sprintf("Received a status update: %s", status)
	// }
}

func (ui *UI) updateStatus(g *gocui.Gui, status string) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View(StatusView)
		if err != nil {
			return err
		}
		v.Clear()
		fmt.Fprintln(v, status)

		return nil
	})
}

func (ui *UI) update(g *gocui.Gui, state state.GameState) {
	g.Update(func(g *gocui.Gui) error {

		// State
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
		ui.showMembers(v, state.MemberStatus)

		// Buffs
		v, err = g.View(BuffsView)
		if err != nil {
			return err
		}
		v.Clear()
		ui.showBuffs(v, state.Buffs)

		// Debuffs
		v, err = g.View(DebuffsView)
		if err != nil {
			return err
		}
		v.Clear()
		// ui.showDebuffs(v, state.Debuffs)
		// ui.showMembers(v, state.MemberStatus)
		ui.showMembers(v, state.MemberCombatStatus)
		ui.showMembers(v, state.MemberMeleeRange)
		ui.show(v, "Dispel", state.Dispel)
		ui.showMembers(v, state.Misc)

		// Target Debuffs
		v, err = g.View(TargetDebuffsView)
		if err != nil {
			return err
		}
		v.Clear()
		ui.showDebuffs(v, state.TargetDebuffs)

		// Spells
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

func (ui *UI) showSpells(v *gocui.View, m map[string]state.SpellData) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		ui.show(v, k, m[k].Keybind)
	}
}

func (ui *UI) showMembers(v *gocui.View, m map[string]bool) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		ui.show(v, k, m[k])
	}
}
