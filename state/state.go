package state

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-vgo/robotgo"
	"github.com/johnayoung/go-wow-bot/utils"
	"github.com/vcaesar/imgo"
)

func getCoordinates() []Cell {
	pathToCoordinates := filepath.Join(utils.RootDir(), "state", "coordinates.json")
	jsonFile, err := os.Open(pathToCoordinates)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Successfully opened coordinates.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var coordinates []Cell
	json.Unmarshal([]byte(byteValue), &coordinates)

	return coordinates
}

func getSpells() Spells {
	pathToSpells := filepath.Join(utils.RootDir(), "state", "spells.json")
	jsonFile, err := os.Open(pathToSpells)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Successfully opened spells.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var spells Spells
	json.Unmarshal([]byte(byteValue), &spells)

	return spells
}

type StateManager struct {
	coordinates []Cell
	spells      Spells
	reader      SquareReader
}

type GameState struct {
	HealthMax,
	HealthCurrent,
	ManaMax,
	ManaCurrent,
	EnergyMax,
	EnergyCurrent,
	Energy,
	ComboPoints,
	RageMax,
	RageCurrent,
	Rage int64
	Health,
	Mana float64
	LowMana       bool
	Buffs         map[string]Buff
	Debuffs       map[string]Debuff
	TargetDebuffs map[string]Debuff
	Spells        map[string]SpellMeta
}

type Pixel struct {
	xMin,
	yMin,
	xMax,
	yMax int
}

func NewStateManager() StateManager {
	return StateManager{
		coordinates: getCoordinates(),
		spells:      getSpells(),
		reader:      NewSquareReader(),
	}
}

func (s StateManager) Get() GameState {
	coordinates := s.coordinates

	pixel := Pixel{
		xMin: 0,
		yMin: 0,
		xMax: 489,
		yMax: 35,
	}
	bit := robotgo.CaptureScreen(pixel.xMin, pixel.yMin, pixel.xMax, pixel.yMax)
	defer robotgo.FreeBitmap(bit)

	// bitmap := robotgo.ToBitmap(bit)

	img := robotgo.ToImage(bit)
	imgo.Save("test.png", img)

	healthMax := s.reader.GetIntAtCell(coordinates[2])
	healthCurrent := s.reader.GetIntAtCell(coordinates[3])
	health := percent(healthCurrent, healthMax)

	// Mana stats
	manaMax := s.reader.GetIntAtCell(coordinates[4])
	manaCurrent := s.reader.GetIntAtCell(coordinates[5])
	mana := percent(manaCurrent, manaMax)
	lowMana := mana < 30

	energyMax := s.reader.GetIntAtCell(coordinates[8])
	energyCurrent := s.reader.GetIntAtCell(coordinates[6])
	comboPoints := s.reader.GetIntAtCell(coordinates[7])

	rageCurrent := s.reader.GetIntAtCell(coordinates[9])
	rageMax := int64(100)

	const numberOfBuffs = 20
	const numberOfDebuffs = 20
	const numberOfAbilities = 48
	startFrame := 11

	buffs := s.GetBuffBlock(startFrame, startFrame+numberOfBuffs)
	startFrame += numberOfBuffs

	debuffs := s.GetDebuffBlock(startFrame, startFrame+numberOfDebuffs)
	startFrame += numberOfDebuffs

	targetDebuffs := s.GetDebuffBlock(startFrame, startFrame+numberOfDebuffs)
	startFrame += numberOfDebuffs

	spells := s.GetSpellBlock(startFrame, startFrame+numberOfAbilities)

	return GameState{
		HealthMax:     healthMax,
		HealthCurrent: healthCurrent,
		Health:        health,
		ManaMax:       manaMax,
		ManaCurrent:   manaCurrent,
		Mana:          mana,
		LowMana:       lowMana,
		EnergyMax:     energyMax,
		EnergyCurrent: energyCurrent,
		ComboPoints:   comboPoints,
		RageMax:       rageMax,
		RageCurrent:   rageCurrent,
		Buffs:         buffs,
		Debuffs:       debuffs,
		TargetDebuffs: targetDebuffs,
		Spells:        spells,
	}
}

type SpellMeta struct {
	spellID   int64
	camelName string
}

type Buff struct {
	SpellMeta
	active bool
}

type Debuff struct {
	SpellMeta
}

func (s StateManager) getBlock(start, end int) []SpellMeta {
	var spells []SpellMeta
	for position := start; position < end; position++ {
		spellID := s.reader.GetIntAtCell(s.coordinates[position])
		spellStr := strconv.FormatInt(spellID, 10)

		if spell, ok := s.spells[spellStr]; ok {
			spells = append(spells, SpellMeta{spellID: spellID, camelName: *spell.CamelName})
		}
	}
	return spells
}

func (s StateManager) GetBuffBlock(
	start,
	end int,
) map[string]Buff {
	var buffs = make(map[string]Buff)

	spells := s.getBlock(start, end)
	for _, spellMeta := range spells {
		buffs[spellMeta.camelName] = Buff{SpellMeta: spellMeta, active: true}
	}

	return buffs
}

func (s StateManager) GetDebuffBlock(
	start,
	end int,
) map[string]Debuff {
	var debuffs = make(map[string]Debuff)

	spells := s.getBlock(start, end)
	for _, spellMeta := range spells {
		debuffs[spellMeta.camelName] = Debuff{SpellMeta: spellMeta}
	}

	return debuffs
}

func (s StateManager) GetSpellBlock(
	start,
	end int,
) map[string]SpellMeta {
	var spellList = make(map[string]SpellMeta)

	spells := s.getBlock(start, end)
	for _, spellMeta := range spells {
		spellList[spellMeta.camelName] = SpellMeta{spellID: spellMeta.spellID, camelName: spellMeta.camelName}
	}

	return spellList
}

func deepCopy(v *string) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	vptr := ""
	err = json.Unmarshal(data, vptr)
	if err != nil {
		return "", err
	}
	return vptr, err
}

func percent(current int64, max int64) float64 {
	diff := float64(current) / float64(max)

	return diff * 100
}
