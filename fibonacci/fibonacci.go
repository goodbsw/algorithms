package fibonacci

// O(2^n) time | O(n) space
func GetNthNumberNaiveVersion(n int) int {
	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	}
	return GetNthNumber(n-1) + GetNthNumber(n-2) // this recursion makes time complexity of O(2^n)
}

// O(n) time | O(n) space
func GetNthNumberMemoizeVersion(n int) int {
	return helper(n, map[int]int{
		1: 0,
		2: 1,
	})
}

func helper(n int, memoize map[int]int) int {
	if val, found := memoize[n]; found {
		return val
	}
	memoize[n] = helper(n-1, memoize) + helper(n-2, memoize)
	return memoize[n]
}

// O(n) time | O(1) space
func GetNthNumber(n int) int {
	lastTwo := []int{0, 1} // makes first two numbers of fibonacci
	counter := 3
	for counter <= n {
		nextFib := lastTwo[0] + lastTwo[1]
		// replace 1st element with 2nd element and 2nd element with nextFib to maintain slice space
		// this makes space complexity of O(1)
		lastTwo[0] = lastTwo[1]
		lastTwo[1] = nextFib
		counter++
	}
	// returns 1st element of lastTwo if n < 1 else returns lastTwo[1]
	if n > 1 {
		return lastTwo[1]
	}
	return lastTwo[0]
}
