//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//Find the sum of all the multiples of 3 or 5 below 1000.
package main

import "fmt"

func main() {
	var sum int
	sum = sum_of_multiples(1000)
	fmt.Println(sum)
}
func sum_of_multiples(x int) int {
	var i, sum int
	sum = 0
	for i = 0; i < x; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
