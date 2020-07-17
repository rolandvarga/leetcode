package main_test

import (
	"math"
	"testing"
)

func minTimeToVisitAllPoints(points [][]int) int {
	minTime := 0

	for i := 0; i < len(points)-1; i++ {
		distX := math.Abs(float64(points[i][0]) - float64(points[i+1][0]))
		distY := math.Abs(float64(points[i][1]) - float64(points[i+1][1]))

		minTime = minTime + int(math.Max(distX, distY))
	}
	return minTime
}

func TestMinTimeToVisitAllPoints(t *testing.T) {
	var cases = []struct {
		points [][]int
		want   int
	}{
		{[][]int{[]int{1, 1}, []int{3, 4}, []int{-1, 0}}, 7},
		{[][]int{[]int{3, 2}, []int{-2, 2}}, 5},
	}

	for _, c := range cases {
		got := minTimeToVisitAllPoints(c.points)
		if got != c.want {
			t.Errorf("got '%d' want '%d'", got, c.want)
		}
	}
}
