package clock

import (
	"math"
	"time"
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
	return math.Pi / (30 / float64(t.Second()))
}

func secondHandPoint(t time.Time) Point {
	radians := secondsInRadians(t)
	return Point{math.Sin(radians), math.Cos(radians)}
}
