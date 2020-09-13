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
	"errors"
	"math/rand"

	u "github.com/csixteen/simulated-evolution/pkg/utils"
)

type Entity interface {
	Id() int
	EntityType() string
	GetPosition() u.Point
	GetEnergy() float64
	Update(w *World)
}

type World struct {
	Width, Height int
	Entities      map[u.Point]Entity
}

func NewWorld(w, h int) *World {
	world := &World{
		Width:    w,
		Height:   h,
		Entities: make(map[u.Point]Entity),
	}

	a := NewAnimal(float64(w/2), float64(h/2))
	world.PlaceEntity(a)

	return world
}

func (w *World) SpawnTree() {
	if rand.Intn(1000) < 150 {
		t := NewTree(
			0,
			0,
			TreeType(rand.Intn(len(treeEnergy))),
		)

		treePlanted := false
		for !treePlanted {
			pos := u.Point{
				float64(rand.Intn(w.Width)),
				float64(rand.Intn(w.Height)),
			}
			if w.IsPlaceVacant(pos) {
				t.Pos = pos
				w.PlaceEntity(t)
				treePlanted = true
			}
		}
	}
}

func (w *World) IsPlaceVacant(p u.Point) bool {
	_, ok := w.Entities[p]
	return !ok
}

func (w *World) PlaceEntity(e Entity) error {
	if !w.IsPlaceVacant(e.GetPosition()) {
		return errors.New("Entity already exists in those coordinates")
	}

	w.Entities[e.GetPosition()] = e

	return nil
}

func (w *World) RemoveEntity(p u.Point) error {
	if w.IsPlaceVacant(p) {
		return errors.New("No entity in those coordinates")
	}

	delete(w.Entities, p)

	return nil
}

func (w *World) UpdateEntities() {
	for p, e := range w.Entities {
		e.Update(w)
		if e.GetEnergy() < 0 {
			w.RemoveEntity(p)
		}
	}
}

func (w *World) Update() {
	w.SpawnTree()
	w.UpdateEntities()
}
