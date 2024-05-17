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
        currentDistance := math.Abs(float64(v))
        closestDistance := math.Abs(float64(closest))
        if currentDistance < closestDistance || (currentDistance == closestDistance && v > 0 && closest < 0) {
            closest = v
        }
    }

    return
}

func newClosest(champion, candidate int) bool {
    return false
}