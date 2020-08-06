package mower

import (
	"testing"

	"github.com/golang/geo/r2"
	"github.com/google/go-cmp/cmp"
	"github.com/valaymerick/mower/lawn"
)

func TestGetAngle(t *testing.T) {
	for i, test := range []struct{ from, angle, want int }{
		// clockwise rotations
		{0, 90, 90},
		{90, 90, 180},
		{180, 90, 270},
		{270, 90, 360},
		{360, 90, 90},

		// anti clockwise rotations
		{0, -90, 270},
		{270, -90, 180},
		{180, -90, 90},
		{90, -90, 0},
	} {
		got := getAngle(test.from, test.angle)
		if got != test.want {
			t.Errorf("test #%d: rotate(%d, %d) = %d; want %d", i, test.from, test.angle, got, test.want)
		}
	}
}

func TestNextPos(t *testing.T) {
	for i, test := range []struct {
		pos   [2]int
		angle int
		move  Move
		want  [2]int
	}{
		{
			pos:   [2]int{0, 0},
			angle: North,
			move:  Forwards,
			want:  [2]int{0, 1},
		},
		{
			pos:   [2]int{0, 1},
			angle: North,
			move:  Backwards,
			want:  [2]int{0, 0},
		},
		{
			pos:   [2]int{0, 0},
			angle: East,
			move:  Forwards,
			want:  [2]int{1, 0},
		},
		{
			pos:   [2]int{1, 0},
			angle: East,
			move:  Backwards,
			want:  [2]int{0, 0},
		},
		{
			pos:   [2]int{0, 0},
			angle: West,
			move:  Forwards,
			want:  [2]int{-1, 0},
		},
		{
			pos:   [2]int{0, 0},
			angle: West,
			move:  Backwards,
			want:  [2]int{1, 0},
		},
		{
			pos:   [2]int{0, 0},
			angle: South,
			move:  Forwards,
			want:  [2]int{0, -1},
		},
		{
			pos:   [2]int{0, 0},
			angle: South,
			move:  Backwards,
			want:  [2]int{0, 1},
		},
	} {
		mower := New(test.pos[0], test.pos[1], test.angle)

		nextPos := mower.nextPos(test.move)

		got := [2]int{int(nextPos.X), int(nextPos.Y)}

		if got != test.want {
			t.Errorf("test #%d: want %d, got %d", i, test.want, got)
		}
	}
}

func TestMow(t *testing.T) {
	l := lawn.New(2, 2)
	m := New(0, 0, North)
	m.Instruct([]Move{
		Right,
		Forwards,
		Left,
		Forwards,
		Left,
		Forwards,
		Left,
		Forwards,
		Backwards,
	})

	m.Mow(&l, nil)

	// Test the final mower position
	want := r2.Point{X: 0, Y: 1}
	got := m.Pos()

	if !cmp.Equal(got, want) {
		t.Errorf("got %+v; want %+v", got, want)
	}
}

func TestMowOutsideLawn(t *testing.T) {
	l := lawn.New(0, 0)
	m := New(0, 0, North)
	m.Instruct([]Move{
		Forwards,
	})

	m.Mow(&l, nil)

	// Check that the mower stops at the Lawn's boundaries
	want := r2.Point{X: 0, Y: 0}
	got := m.Pos()

	if !cmp.Equal(got, want) {
		t.Errorf("got %+v; want %+v", got, want)
	}
}
