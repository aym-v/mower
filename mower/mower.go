package mower

import (
	"fmt"
	"sync"

	"github.com/golang/geo/r2"
	geo "github.com/golang/geo/r2"
	"github.com/valaymerick/mower/lawn"
)

// Orientation represents a cardinal direction.
type Orientation int

// Move represents a single-axis rotation or translation.
type Move int

// Cardinal angles
const (
	North = 0
	East  = 90
	South = 180
	West  = 270
)

// Mower movements.
const (
	Left Move = iota
	Right
	Forwards
	Backwards
)

// Mower represents a mower automata.
type Mower struct {
	Position     geo.Point
	Orientation  int
	Instructions []Move
}

// New returns a new mower at the corresponding location.
func New(x, y int, orientation int) Mower {
	instructions := make([]Move, 0)
	position := r2.Point{X: float64(x), Y: float64(y)}

	return Mower{
		position,
		orientation,
		instructions,
	}
}

// Instruct adds the given instructions to the mower's state.
func (m *Mower) Instruct(moves []Move) {
	m.Instructions = moves
}

// Mow mows the given lawn according to the instructions.
func (m *Mower) Mow(lawn *lawn.Lawn, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	for _, instr := range m.Instructions {
		m.move(instr, lawn)
	}
}

// move executes a move of the given type.
func (m *Mower) move(typ Move, lawn *lawn.Lawn) {
	switch typ {
	case Left:
		m.rotate(-90)
	case Right:
		m.rotate(90)

	case Forwards:
		n := m.nextPos(Forwards)

		if lawn.IsMowable(n) {
			m.Position = n
			lawn.Release(n)
		}

	case Backwards:
		n := m.nextPos(Backwards)

		if lawn.IsMowable(n) {
			m.Position = n
			lawn.Release(n)
		}
	}
}

// rotate performs a rotation around the Z axis.
func (m *Mower) rotate(angle int) {
	m.Orientation = getAngle(m.Orientation, angle)
}

// nextPos returns the position that the mower will have if it performs the given movement.
func (m *Mower) nextPos(flow Move) geo.Point {
	x, y := 0.0, 0.0
	dir := 1.0

	if flow == Backwards {
		dir = -1
	}

	switch m.Orientation {
	case North:
		y = dir * 1
	case East:
		x = dir * 1
	case South:
		y = dir * -1
	case West:
		x = dir * -1
	}

	return geo.Point{X: m.Position.X + x, Y: m.Position.Y + y}
}

// String returns a mower string representation in the form `X Y OrientationKey`.
func (m *Mower) String() string {
	return fmt.Sprintf("%v %v %v", int(m.Position.X), int(m.Position.Y), getOrientationKey(m.Orientation))
}

// Pos returns the position of the mower.
func (m *Mower) Pos() geo.Point {
	return m.Position
}

// getAngle returns the orientation resulting from a rotation of `angle` degrees
func getAngle(from int, angle int) int {
	r := from + angle

	if r > 360 {
		r -= 360
	} else if r < 0 {
		r += 360
	}

	return r
}

// getOrientationKey maps an angle to a configuration key
func getOrientationKey(angle int) string {
	switch angle {
	case 0:
		return "N"
	case 90:
		return "E"
	case 180:
		return "S"
	case 270:
		return "W"
	default:
		return ""
	}
}
