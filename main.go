package main

import (
	"fmt"
)

func main() {
	data := []int{200, 100}

	sum := sumData(data)
	fmt.Println("Sum:", sum)
	fmt.Println("To Sum One Digit:", sumOneDigit(sum))
	// fmt.Println(300 % 10) = 0
	// fmt.Println(300 / 10) = 30
	// 0 + 30
	// fmt.Println(30 % 10) = 0
	// fmt.Println(30 / 10) = 3
	// 0 + 3
	// fmt.Println(3 % 10) = 3
	// fmt.Println(3 / 10) = 0

}

func sumData(data []int) int {
	sum := 0
	for i := 0; i < len(data); i++ {
		sum += data[i]
	}
	return sum
}

func sumOneDigit(n int) int {
	if n == 0 {
		return n
	}
	return n%10 + sumOneDigit(n/10)
}
