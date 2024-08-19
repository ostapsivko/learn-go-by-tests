package clock

import (
	"math"
	"time"
)

type Point struct {
	X, Y float64
}

func SecondHand(date time.Time) Point {
	return Point{150, 60}
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}
