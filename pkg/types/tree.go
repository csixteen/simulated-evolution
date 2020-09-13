package types

import (
	u "github.com/csixteen/simulated-evolution/pkg/utils"
)

type TreeType int32

const (
	Poisonous  TreeType = 0
	Neutral1   TreeType = 1
	Nutritious TreeType = 2
	Neutral2   TreeType = 3
	Energetic  TreeType = 4
	Neutral3   TreeType = 5
)

var treeEnergy = map[TreeType]float64{
	Poisonous:  -300,
	Neutral1:   10,
	Nutritious: 500,
	Neutral2:   15,
	Energetic:  1000,
	Neutral3:   20,
}

type Tree struct {
	Type   TreeType
	Pos    u.Point
	Energy float64
}

func NewTree(x, y float64, t TreeType) *Tree {
	e, _ := treeEnergy[t] // Check for error, perhaps?
	return &Tree{
		Pos:    u.Point{x, y},
		Energy: e,
		Type:   t,
	}
}

func (t *Tree) GetPosition() u.Point {
	return t.Pos
}

func (t *Tree) Update(world *World) {
}

func (t *Tree) EntityType() string {
	return "tree"
}

func (t *Tree) Id() int {
	return int(t.Type)
}

func (t *Tree) GetEnergy() float64 {
	return t.Energy
}
