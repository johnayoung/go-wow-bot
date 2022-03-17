package state

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-vgo/robotgo"
	"github.com/johnayoung/go-wow-bot/utils"
	"github.com/vcaesar/imgo"
)

type Member string

const (
	Target    = "target"
	Focus     = "focus"
	Mouseover = "mouseover"
	Player    = "player"
	Pet       = "pet"
	Party1    = "party1"
	Party2    = "party2"
	Party3    = "party3"
	Party4    = "party4"
	Party5    = "party5"
)

var listOfMembers = []string{
	Target,
	Focus,
	Mouseover,
	Player,
	Pet,
	Party1,
	Party2,
	Party3,
	Party4,
	Party5,
}

var miscValue = []string{
	"hasWeaponEnchant",
	"maelstromWeapon",
	"wildWrath",
	"isFrozen",
	"targetPurgeable",
}

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
	Rage,
	Dispel int64
	Health,
	Mana float64
	LowMana            bool
	Buffs              map[string]Buff
	Debuffs            map[string]Debuff
	TargetDebuffs      map[string]Debuff
	Spells             map[string]SpellData
	MemberStatus       map[string]bool
	MemberCombatStatus map[string]bool
	MemberMeleeRange   map[string]bool
	Misc               map[string]bool
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
	startFrame += numberOfAbilities

	memberStatus := s.reader.GetIntAtCell(coordinates[startFrame])
	startFrame += 1

	statusCoords := coordinates[startFrame]
	memberCombatStatus := s.reader.GetIntAtCell(statusCoords)
	startFrame += 1
	memberMeleeRange := s.reader.GetIntAtCell(coordinates[startFrame])
	startFrame += 1
	dispel := s.reader.GetIntAtCell(coordinates[startFrame])
	startFrame += 1

	miscBinary := s.reader.GetIntAtCell(coordinates[startFrame])
	// misList :=
	startFrame += 1

	return GameState{
		HealthMax:          healthMax,
		HealthCurrent:      healthCurrent,
		Health:             health,
		ManaMax:            manaMax,
		ManaCurrent:        manaCurrent,
		Mana:               mana,
		LowMana:            lowMana,
		EnergyMax:          energyMax,
		EnergyCurrent:      energyCurrent,
		ComboPoints:        comboPoints,
		RageMax:            rageMax,
		RageCurrent:        rageCurrent,
		Buffs:              buffs,
		Debuffs:            debuffs,
		TargetDebuffs:      targetDebuffs,
		Spells:             spells,
		MemberStatus:       s.GetBoolStatus(memberStatus, listOfMembers, "IsDead"),
		MemberCombatStatus: s.GetBoolStatus(memberCombatStatus, listOfMembers, "InCombat"),
		MemberMeleeRange:   s.GetBoolStatus(memberMeleeRange, listOfMembers, "InMeleeRange"),
		Misc:               s.GetBoolStatus(miscBinary, miscValue, ""),
		Dispel:             dispel,
	}
}

type SpellMeta struct {
	spellID   int64
	camelName string
	counter   int
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
		counter := position - start
		spellID := s.reader.GetIntAtCell(s.coordinates[position])
		spellStr := strconv.FormatInt(spellID, 10)

		if spell, ok := s.spells[spellStr]; ok {
			spells = append(spells, SpellMeta{spellID: spellID, camelName: *spell.CamelName, counter: counter})
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

var keyBindings = map[int]string{
	// Action bar 1
	0:  "g",
	1:  "r",
	2:  "t",
	3:  "v",
	4:  "w",
	5:  "a",
	6:  "q",
	7:  "z",
	8:  "c",
	9:  "4",
	10: "5",
	11: "6",
	12: "y",
	13: "h",
	14: "b",
	15: "1",
	16: "2",
	17: "3",
	18: "7",
	19: "8",
	20: "9",
	21: "0",
	22: "",
	23: "",
	// Bottom Right ActionBar
	24: "f1",
	25: "f2",
	26: "f3",
	27: "f4",
	28: "f5",
	29: "f6",
	30: "f7",
	31: "f8",
	32: "f9",
	33: "f10",
	34: "f11",
	35: "f12",
	// Shapeshift Bar
	36: "g",
	37: "r",
	38: "t",
	39: "v",
	40: "w",
	41: "a",
	42: "q",
	43: "z",
	44: "c",
	45: "4",
	46: "5",
	47: "6",
}

type SpellData struct {
	SpellMeta
	keybind  string
	castable bool
	equipped bool
}

func (s StateManager) GetSpellBlock(
	start,
	end int,
) map[string]SpellData {
	var spellList = make(map[string]SpellData)

	spells := s.getBlock(start, end)
	for _, spellMeta := range spells {
		meta := SpellMeta{spellID: spellMeta.spellID, camelName: spellMeta.camelName}
		spellList[spellMeta.camelName] = SpellData{
			spellMeta,
			keyBindings[meta.counter],
			true,
			true,
		}
	}

	return spellList
}

func (s StateManager) GetBoolStatus(base2Operation int64, list []string, suffix string) map[string]bool {
	var members = make(map[string]bool)
	var memberCount = len(list)

	for i := memberCount; i > 0; i-- {
		decrease := int64(math.Pow(2, float64(i)))
		active := base2Operation-decrease >= 0

		unit := list[i-1]
		memberKey := fmt.Sprintf("%s%s", unit, suffix)
		members[memberKey] = active

		if active == true {
			base2Operation -= decrease
		}
	}

	return members
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
