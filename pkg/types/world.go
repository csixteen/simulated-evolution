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
