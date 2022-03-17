// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    spells, err := UnmarshalSpells(bytes)
//    bytes, err = spells.Marshal()

package state

import "encoding/json"

type Spells map[string]Spell

func UnmarshalSpells(data []byte) (Spells, error) {
	var r Spells
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Spells) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Spell struct {
	SpellID          *int64            `json:"id,omitempty"`
	Name             *string           `json:"name,omitempty"`
	ImageURL         *string           `json:"imageUrl,omitempty"`
	Type             *string           `json:"type,omitempty"`
	Ranks            *Ranks            `json:"ranks,omitempty"`
	CastTime         *CastTime         `json:"castTime,omitempty"`
	Cooldown         *CastTime         `json:"cooldown,omitempty"`
	Tooltip          *Tooltip          `json:"tooltip,omitempty"`
	LearnableAbility *LearnableAbility `json:"learnableAbility,omitempty"`
	ID               *string           `json:"_id,omitempty"`
	CamelName        *string           `json:"camelName,omitempty"`
	Cost             *Cost             `json:"cost,omitempty"`
	Range            *Range            `json:"range,omitempty"`
	Duration         *CastTime         `json:"duration,omitempty"`
	LearnableTalent  *LearnableTalent  `json:"learnableTalent,omitempty"`
}

type CastTime struct {
	Type    *string  `json:"type,omitempty"`
	Text    *string  `json:"text,omitempty"`
	Seconds *float64 `json:"seconds,omitempty"`
}

type Cost struct {
	Type  *string `json:"type,omitempty"`
	Value *int64  `json:"value,omitempty"`
	Text  *string `json:"text,omitempty"`
}

type LearnableAbility struct {
	RequiredLevel      *int64 `json:"requiredLevel,omitempty"`
	AbilityEssenceCost *int64 `json:"abilityEssenceCost,omitempty"`
	ClassID            *int64 `json:"classId,omitempty"`
	SpecializationID   *int64 `json:"specializationId,omitempty"`
}

type LearnableTalent struct {
	RequiredLevel      *int64 `json:"requiredLevel,omitempty"`
	AbilityEssenceCost *int64 `json:"abilityEssenceCost,omitempty"`
	TalentEssenceCost  *int64 `json:"talentEssenceCost,omitempty"`
	TalentID           *int64 `json:"talentId,omitempty"`
	ClassID            *int64 `json:"classId,omitempty"`
	SpecializationID   *int64 `json:"specializationId,omitempty"`
	VisualColumn       *int64 `json:"visualColumn,omitempty"`
	MaxRank            *int64 `json:"maxRank,omitempty"`
}

type Range struct {
	Yards *int64  `json:"yards,omitempty"`
	Type  *string `json:"type,omitempty"`
	Text  *string `json:"text,omitempty"`
}

type Ranks struct {
	Current *int64  `json:"current,omitempty"`
	Other   []int64 `json:"other,omitempty"`
	Text    *string `json:"text,omitempty"`
}

type Tooltip struct {
	Text       *string `json:"text,omitempty"`
	TooltipURL *string `json:"tooltipUrl,omitempty"`
}
