// MIT License
//
// Copyright (c) 2020 Pedro Rodrigues
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

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
