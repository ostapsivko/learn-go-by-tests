package clock

import (
	"math"
	"time"
)

// TODO: reorganize code and document all exported functions

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

type Point struct {
	X, Y float64
}

// SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point.
func SecondHand(date time.Time) Point {
	p := secondHandPoint(date)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} //scale
	p = Point{p.X, -p.Y}                                      //flip
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}         //translate
	return p
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (secondsInHalfClock / float64(t.Second()))
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / secondsInClock) +
		math.Pi/(minutesInHalfClock/float64(t.Minute()))
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / hoursInClock) +
		math.Pi/(hoursInHalfClock/float64(t.Hour()%hoursInClock))
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
