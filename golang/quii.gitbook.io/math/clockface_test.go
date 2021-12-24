package math

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 7), math.Pi / 30 * 7},
		{simpleTime(0, 0, 15), math.Pi / 2},
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 45), math.Pi * 1.5},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondsInRadians(c.time)
			if !routhlyEqualFloat64(got, c.angle) {
				t.Fatalf("Wanted %g radians, but got %g", c.angle, got)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
		{simpleTime(0, 15, 0), math.Pi / 2},
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 45, 0), math.Pi * 1.5},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := minutesInRadians(c.time)
			if !routhlyEqualFloat64(got, c.angle) {
				t.Fatalf("Wanted %g radians, but got %g", c.angle, got)
			}
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
		{simpleTime(3, 0, 0), math.Pi / 2},
		{simpleTime(18, 0, 0), math.Pi},
		{simpleTime(21, 0, 0), math.Pi * 1.5},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := hoursInRadians(c.time)
			if !routhlyEqualFloat64(got, c.angle) {
				t.Fatalf("Wanted %g radians, but got %g", c.angle, got)
			}
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 15), Point{1, 0}},
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
		{simpleTime(0, 0, 60), Point{0, 1}},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondHandPoint(c.time)
			if !routhlyEqualPoint(got, c.point) {
				t.Fatalf("wanted %+v Point, but got %+v", c.point, got)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 15, 0), Point{1, 0}},
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
		{simpleTime(0, 60, 0), Point{0, 1}},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := minuteHandPoint(c.time)
			if !routhlyEqualPoint(got, c.point) {
				t.Fatalf("wanted %+v Point, but got %+v", c.point, got)
			}
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(3, 0, 0), Point{1, 0}},
		{simpleTime(18, 0, 0), Point{0, -1}},
		{simpleTime(21, 0, 0), Point{-1, 0}},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := hourHandPoint(c.time)
			if !routhlyEqualPoint(got, c.point) {
				t.Fatalf("wanted %+v Point, but got %+v", c.point, got)
			}
		})
	}
}

func routhlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func routhlyEqualPoint(a, b Point) bool {
	return routhlyEqualFloat64(a.X, b.X) && routhlyEqualFloat64(a.Y, b.Y)
}

func simpleTime(h, m, s int) time.Time {
	return time.Date(312, time.October, 28, h, m, s, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
