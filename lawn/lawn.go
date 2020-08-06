package lawn

import (
	"sync"

	"github.com/golang/geo/r1"
	"github.com/golang/geo/r2"
	"golang.org/x/sync/semaphore"
)

var mutex = sync.RWMutex{}

// Lawn represents a lawn composed of Plots
type Lawn struct {
	// Area represents the ground dimensions
	Area r2.Rect
	// Plots represents each mowable plot
	Plots map[r2.Point]*semaphore.Weighted
}

// New creates a new virtual lawn with the given size.
func New(x int, y int) Lawn {
	surface := r2.Rect{
		X: r1.Interval{Lo: 0, Hi: float64(x)},
		Y: r1.Interval{Lo: 0, Hi: float64(y)},
	}

	return Lawn{
		Area:  surface,
		Plots: make(map[r2.Point]*semaphore.Weighted),
	}
}

// acquire tries to acquire the given plot.
// On success, returns true. On failure, returns false
func (l *Lawn) acquire(plot r2.Point) bool {
	mutex.RLock()
	_, ok := l.Plots[plot]
	mutex.RUnlock()

	// If no semaphore exists for this key, create it
	if !ok {
		mutex.Lock()
		l.Plots[plot] = semaphore.NewWeighted(1)
		mutex.Unlock()
	}

	return l.Plots[plot].TryAcquire(1)
}

// Release releases the given plot.
func (l *Lawn) Release(plot r2.Point) {
	l.Plots[plot].Release(1)
}

// IsMowable reports whether the given plot is available for mowing.
func (l *Lawn) IsMowable(plot r2.Point) bool {
	return l.acquire(plot) && l.Area.ContainsPoint(plot)
}
