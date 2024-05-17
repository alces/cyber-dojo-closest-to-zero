package ctz

import (
    "fmt"
    "testing"

    "github.com/stretchr/testify/assert"
)

var closestToZeroTestResults = []struct {
    argument []int
    expected int
} {
    {[]int{}, 0},
    {[]int{2}, 2},
    {[]int{2, 1}, 1},
    {[]int{2, -1, 3}, -1},
    {[]int{1, 2, -1, 3}, 1},
    {[]int{-1, 2, 1, 3}, 1},
}

var newClosestTestResults = []struct {
    champion  int
    candidate int
    expected bool
} {
    {2, 1, true},
    {2, 3, false},
    {2, -2, false},
    {-2, 2, false},
}

func TestClosestToZero(t *testing.T) {
    for _, res := range closestToZeroTestResults {
        assert.Equal(t, res.expected, closestToZero(res.argument), fmt.Sprintf("slice: %#v", res.argument))
    }
}

func TestNewClosest(t *testing.T) {
for _, res := range newClosestTestResults {
        assert.Equal(t, res.expected, newClosest(res.champion, res.candidate), fmt.Sprintf("champion: %d candidate: %d", res.champion, res.candidate))
    }
}