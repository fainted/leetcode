package main

import (
	"log"
	"math"
)

func reverse(x int) int {
	absValue := x
	if absValue < 0 {
		absValue = -absValue
	}

	var result int
	for absValue != 0 {
		temp := result*10 + absValue%10
		absValue /= 10
		if temp < result {
			return 0 // overflow
		}
		result = temp
	}

	if x >= 0 && result <= math.MaxInt32 {
		return result
	} else if x < 0 && -result >= math.MinInt32 {
		return -result
	}
	return 0
}

func main() {
	log.Print(reverse(7463847412))
}
