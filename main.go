package main

import (
	"fmt"

	"github.com/goodbsw/algorithms/twonumbersum"
)

func main() {
	// fmt.Println(fibonacci.GetNthNumberNaiveVersion(5))
	// fmt.Println(fibonacci.GetNthNumberMemoizeVersion(5))
	// fmt.Println(fibonacci.GetNthNumber(5))
	fmt.Println(twonumbersum.TwoNumberSumWithNaiveVersion([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10))
	fmt.Println(twonumbersum.TwoNumberSumWithHashVersion([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10))
	fmt.Println(twonumbersum.TwoNumberSumWithSpaceOptimalVersion([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10))
}
