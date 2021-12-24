package math

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

// A Point represents a two dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

func angleToPoint(angle float64) Point {
	return Point{
		X: math.Sin(angle),
		Y: math.Cos(angle),
	}
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (secondsInHalfClock / float64(t.Second()))
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / minutesInClock) + (math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / hoursInClock) + math.Pi/(hoursInHalfClock/(float64(t.Hour()%hoursInClock)))
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}
