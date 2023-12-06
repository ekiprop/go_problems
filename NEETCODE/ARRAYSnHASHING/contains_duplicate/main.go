package main

import "fmt"

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		// If the element is already in the map, it's a duplicate
		if seen[num] {
			return true
		}

		// Otherwise, mark the element as seen
		seen[num] = true
	}

	// No duplicates found
	return false
}

func main() {
	// Example usage
	nums1 := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums1)) // Output: true

	nums2 := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums2)) // Output: false

	nums3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(nums3)) // Output: true
}

//This solution also has a time complexity of O(n) and a space complexity of O(n).

// BRUTE FORCE
// Time: O(n*n) = O(n^2)
// Space: O(1)
//---------//
// func containsDuplicate(nums []int) bool {
// 	// Time: O(n-1) = O(n)
// 	for i := 0; i < len(nums)-1; i++ {
// 		// Time: O(n-1) = O(n)
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i] == nums[j] {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
