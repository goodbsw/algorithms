package main

import (
	"fmt"

	"github.com/goodbsw/algorithms/fibonacci"
)

func main() {
	fmt.Println(fibonacci.GetNthNumberNaiveVersion(5))
	fmt.Println(fibonacci.GetNthNumberMemoizeVersion(5))
	fmt.Println(fibonacci.GetNthNumber(5))
}
