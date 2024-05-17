package ctz

import (
    "math"
)

func closestToZero(nums []int) (closest int) {
    if len(nums) == 0 {
        return
    }
    
    closest = nums[0]
    
    for _, v := range nums[1:] {
        if newClosest(closest, v) {
            closest = v
        }
    }

    return
}

func newClosest(champion, candidate int) bool {
    currentDistance := math.Abs(float64(candidate))
    closestDistance := math.Abs(float64(champion))

    if currentDistance < closestDistance {
        return true
    }
    
    return currentDistance == closestDistance && candidate > 0 && champion < 0
}