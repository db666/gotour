// Find the element in an odd length array which is not repeated
package main

import "fmt"

func OddOccur(arr []int) int {
	n := 0
	for _, num := range arr {
		n = n ^ num
	}
	return n
}

func main() {
	n := OddOccur([]int{9,3,9,3,9,7,9})
	fmt.Println("odd number is ", n)
}