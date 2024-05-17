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
        if math.Abs(float64(v)) < float64(closest) {
            closest = v
        }
    }

    return
}
