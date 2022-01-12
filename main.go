package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("hello")
	// s := strconv.QuoteRuneToASCII(0x85)

	fmt.Println(fmt.Sprintf("%c", 0x85))

	// n := 10
	// numbers := []int{}
	// for i := 0; i < n; i++ {
	// 	r := generateNumber()
	// 	numbers = append(numbers, r)
	// }

	// fmt.Println(numbers)
	// fmt.Println(average(numbers))
	// fmt.Println(highestNumber(numbers))
}

func generateNumber() int {
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(100)

	return result
}

func highestNumber(numbers []int) int {
	ans := -1

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > ans {
			ans = numbers[i]
		}
	}

	return ans
}

func average(numbers []int) float64 {
	var total float64
	n := float64(len(numbers))
	for i := 0; i < len(numbers); i++ {
		total += float64(numbers[i])
	}

	return total / n

}
