package ctz

import (
    "fmt"
    "testing"

    "github.com/stretchr/testify/assert"
)

var testResults = []struct {
    argument []int
    expected int
} {
    {[]int{}, 0},
}

func TestEmptySlice(t *testing.T) {
    for _, res := range testResults {
        assert.Equal(t, res.expected, closestToZero(res.argument), fmt.Sprintf("slice: %#v", res.argument))
    }
}
