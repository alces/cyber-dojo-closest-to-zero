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

func newClosest(champion, candidate int) (answer bool) {
    currentDistance := math.Abs(float64(candidate))
    closestDistance := math.Abs(float64(champion))

    if currentDistance < closestDistance {
        answer = true
    } else if currentDistance == closestDistance {
        answer = candidate > 0 && champion < 0
    }
     
    return
}