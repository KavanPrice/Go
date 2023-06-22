package main

import (
	"errors"
	"fmt"
)

func main() {
	intSlice := func(N int) []int {
		slice := make([]int, N+1)
		for i := 0; i < N+1; i++ {
			slice[i] = i
		}
		return slice
	}(10)

	for index := range intSlice {
		if intSlice[index]%2 == 0 {
			fmt.Println(index, "is even")
		} else if intSlice[index]%2 == 1 {
			fmt.Println(index, "is odd")
		} else {
			fmt.Println(errors.New("error occured computing modular division"))
		}
	}
}
