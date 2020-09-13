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
	"math"
	"math/rand"

	u "github.com/csixteen/simulated-evolution/pkg/utils"
)

type Direction int32

const (
	C  Direction = 0
	N  Direction = 1
	NE Direction = 2
	E  Direction = 3
	SE Direction = 4
	S  Direction = 5
	SW Direction = 6
	W  Direction = 7
	NW Direction = 8
)

const reproducingEnergy = 2000

type Animal struct {
	id     int
	Pos    u.Point
	Energy float64
	Dir    Direction
	Genes  [8]float64
}

func (a *Animal) GetPosition() u.Point {
	return a.Pos
}

func (a *Animal) EntityType() string {
	return "animal"
}

func NewAnimal(x, y float64) *Animal {
	a := &Animal{
		id:     rand.Int(),
		Pos:    u.Point{x, y},
		Energy: 10000,
		Dir:    C,
	}

	for i := 0; i < len(a.Genes); i++ {
		a.Genes[i] = float64(rand.Intn(1000))
	}

	return a
}

func (parent *Animal) Reproduce() *Animal {
	prob := rand.Intn(10000)

	if prob < 10 && parent.Energy >= reproducingEnergy {
		parent.Energy /= 2

		child := NewAnimal(
			parent.Pos.X+float64(rand.Intn(50)),
			parent.Pos.Y+float64(rand.Intn(50)),
		)
		child.Genes = parent.Genes
		mutation := rand.Intn(8)
		child.Genes[mutation] = float64(rand.Intn(5000))

		return child
	}

	return nil
}

func (a *Animal) MaybeKill(o *Animal, world *World) {
	// One of them is agressive enough
	if a.Genes[0] > 500 || o.Genes[0] > 500 {
		newEnergy := a.Energy*a.Genes[0]*a.Genes[1] - o.Energy*o.Genes[0]*o.Genes[1]

		a.Energy += newEnergy
		o.Energy += -newEnergy

		if newEnergy > 0 {
			a.Genes[0] += 100
			o.Genes[0] = math.Max(0, o.Genes[0]-100)
		} else {
			a.Genes[0] = math.Max(0, a.Genes[0]-100)
			o.Genes[0] += 100
		}
	}
}

func (a *Animal) Eat(other Entity, world *World) {
	energy := other.GetEnergy()
	a.Energy += energy
	a.Genes[0] = math.Max(0, a.Genes[0]-10)
	world.RemoveEntity(other.GetPosition())
}

func (a *Animal) Interact(other Entity, world *World) {
	switch other.EntityType() {
	case "tree":
		a.Eat(other, world)
	case "animal":
		a.MaybeKill(other.(*Animal), world)
	}
}

func (a *Animal) Explore(world *World) {
	for x := a.Pos.X - 24; x <= a.Pos.X+24; x += 1 {
		for y := a.Pos.Y - 24; y <= a.Pos.Y+24; y += 1 {
			if x != a.Pos.X && y != a.Pos.Y {
				p := u.Point{x, y}
				if !world.IsPlaceVacant(p) {
					a.Interact(world.Entities[p], world)
				}
			}
		}
	}
}

func (a *Animal) Turn() {
	a.Dir = Direction(rand.Intn(9))
}

func (a *Animal) Move(world *World) {
	world.RemoveEntity(a.Pos)

	x := a.Pos.X
	y := a.Pos.Y

	switch a.Dir {
	case N, NE, NW:
		y += 1
	case S, SE, SW:
		y -= 1
	}

	switch a.Dir {
	case SE, E, NE:
		x += 1
	case SW, W, NW:
		x -= 1
	}

	a.Energy -= 0.01
	a.Genes[0] += 0.1 // The more it walks, the more agressive it becomes
	a.Pos.X = x
	a.Pos.Y = y

	world.PlaceEntity(a)
}

func (a *Animal) Update(world *World) {
	a.Explore(world)
	a.Turn()
	a.Move(world)

	child := a.Reproduce()
	if child != nil {
		world.PlaceEntity(child)
	}
}

func (a *Animal) Id() int {
	return a.id
}

func (a *Animal) GetEnergy() float64 {
	return a.Energy
}
