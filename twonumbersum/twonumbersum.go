package twonumbersum

import "sort"

// O(2^N) time | O(1) space
func TwoNumberSumWithNaiveVersion(array []int, target int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array)-1; j++ {
			if array[i]+array[j] == target {
				return []int{array[i], array[j]}
			}
		}
	}
	return nil
}

// O(N) time | O(N) space
func TwoNumberSumWithHashVersion(array []int, target int) []int {
	nums := map[int]bool{} //Create empty map
	for _, num := range array {
		potentialMatch := target - num // Using x + y = target, x=num, y=potentialMatch
		if _, found := nums[potentialMatch]; found {
			return []int{num, potentialMatch}
		}
		nums[num] = true
	}
	return nil
}

// O(log(N)) time | O(1) space
func TwoNumberSumWithSpaceOptimalVersion(array []int, target int) []int {
	sort.Ints(array) //sorting array with increasing order
	right := len(array) - 1
	left := 0

	for left < right { // While left and right are not overlapped
		currentSum := array[left] + array[right]
		if currentSum == target {
			return []int{array[left], array[right]}
		} else if currentSum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}
