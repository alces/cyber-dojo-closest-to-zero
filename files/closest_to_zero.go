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
        distance := int(math.Abs(float64(v)))
        if distance < closest || (distance == closest && v > 0 && closest < 0) {
            closest = v
        }
    }

    return
}
