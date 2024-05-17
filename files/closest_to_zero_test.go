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

func TestClosestToZero(t *testing.T) {
    for _, res := range closestToZeroTestResults {
        assert.Equal(t, res.expected, closestToZero(res.argument), fmt.Sprintf("slice: %#v", res.argument))
    }
}
