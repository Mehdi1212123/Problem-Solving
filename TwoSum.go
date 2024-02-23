package main

import (
	"fmt"
)

// Your twoSum function goes here
func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)
    
    for i, num := range nums {
        complement := target - num
        if j, ok := numMap[complement]; ok {
            return []int{j, i}
        }
        numMap[num] = i
    }
    
    return nil
}

func main() {
    // Example test data
    nums := []int{2,  7,  11,  15}
    target :=  9

    // Call the twoSum function with the test data
    result := twoSum(nums, target)

    // Check if the result is not nil (meaning a pair was found)
    if result != nil {
        fmt.Printf("Indices of the two numbers that add up to %d are %v\n", target, result)
    } else {
        fmt.Println("No pair found that adds up to the target.")
    }
}

