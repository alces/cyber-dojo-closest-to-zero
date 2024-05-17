package ctz

func closestToZero(nums []int) (closest int) {
    if len(nums) == 0 {
        return
    }
    
    closest = nums[0]
    
    for _, v := range nums[1:] {
        if abs(v) < closest {
            closest = v
        }
    }

    return
}
