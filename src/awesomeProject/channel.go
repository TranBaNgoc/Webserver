package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0,1,2,4,5,8,7,5,2,4,5,6,3,2,1,5,4,5,5,4,5,2,1,2,2,2,1,7,5,-75,4}

	c := make(chan int)

	go sum(s[:len(s)/2], c)

	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
