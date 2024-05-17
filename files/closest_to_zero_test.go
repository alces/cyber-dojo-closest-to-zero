package ctz

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestEmptySlice(t *testing.T) {
    assert.Equal(t, 0, closestToZero([]int{}))
}
