package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// Số th
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else {
		return math.Sqrt(x), nil
	}

}

func main() {
	if result, err := Sqrt(2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	if result, err := Sqrt(-2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
