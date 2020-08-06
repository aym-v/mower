package lawn

import (
	"testing"

	"github.com/golang/geo/r2"
)

func TestIsMowable(t *testing.T) {
	for i, test := range []struct {
		lawnWidth, lawnHeight int
		point                 r2.Point
		want                  bool
	}{
		// Inside bounding box
		{5, 5, r2.Point{X: 0.0, Y: 0.0}, true},

		// Outside bounding box
		{5, 5, r2.Point{X: 0.0, Y: 6.0}, false},
	} {
		l := New(test.lawnHeight, test.lawnHeight)

		got := l.IsMowable(test.point)
		if got != test.want {
			t.Errorf("test #%d: got %v; want %v", i, got, test.want)
		}
	}
}

func TestIsMowableLock(t *testing.T) {
	test := struct {
		lawnWidth, lawnHeight int
		plot                  r2.Point
	}{
		5, 5, r2.Point{X: 0.0, Y: 0.0},
	}

	l := New(test.lawnHeight, test.lawnHeight)

	// Acquire a plot and try to acquire it again
	l.acquire(test.plot)
	got := l.acquire(test.plot)

	if got == true {
		t.Errorf("Locked the plot, got %v; want %v", got, false)
	}

	// Release the plot and try to acquire it
	l.Release(test.plot)
	got = l.acquire(test.plot)

	if got == false {
		t.Errorf("Released the plot, got %v; want %v", got, true)
	}
}
