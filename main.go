package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/johnayoung/go-wow-bot/run"
	"github.com/johnayoung/go-wow-bot/state"
	"github.com/johnayoung/go-wow-bot/ui"
	"github.com/jroimartin/gocui"
)

const (
	POLL_INTERVAL_MS = 300
)

func main() {
	updateTicker := time.Tick(POLL_INTERVAL_MS * time.Millisecond)

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	stateView := NewLabel(ui.StateView, g)
	buffsView := NewLabel(ui.BuffsView, g)
	spellView := NewLabel(ui.SpellView, g)
	fl := gocui.ManagerFunc(flowLayout)
	g.SetManager(stateView, buffsView, spellView, fl)

	stateResults := make(chan state.GameState)
	errors := make(chan error)

	agent := &run.Agent{
		Status:       "starting",
		UpdateTicker: updateTicker,
		Errors:       errors,
		StateManager: state.NewStateManager(),
		StateResults: stateResults,
	}

	ui := &ui.UI{
		Errors:       errors,
		StateResults: stateResults,
	}

	go agent.Start()
	go ui.Start(g)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

	for err := range errors {
		log.Fatal(err)
	}
}

type Label struct {
	name string
	w, h int
}

func NewLabel(name string, g *gocui.Gui) *Label {
	maxX, maxY := g.Size()

	w := maxX / 5
	h := maxY - 1

	return &Label{name: name, w: w, h: h}
}

func (w *Label) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, 0, 0, w.w, w.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprint(v, "")
	}
	v.Title = (w.name)
	return nil
}

func flowLayout(g *gocui.Gui) error {
	views := g.Views()
	x := 0
	for _, v := range views {
		w, h := v.Size()
		_, err := g.SetView(v.Name(), x, 0, x+w+1, h+1)
		if err != nil && err != gocui.ErrUnknownView {
			return err
		}
		x += w + 2
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	e := gocui.ErrQuit
	os.Exit(0)
	return e
}
