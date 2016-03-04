package main

import "fmt"

// Rotate array by K positions and return the newly created array

func Rotate(A []int, K int) []int {
    alen := len(A)
    newarr := make([]int, alen)
    var newpos int
    for pos, e := range A {
        newpos = (pos + K) % alen
        newarr[newpos] = e
    }
    return newarr
}

func main() {
	arr := []int{3, 8, 9, 7, 6}
	cnt := 3
	fmt.Println(Rotate(arr, cnt)) // [9, 7, 6, 3, 8]
}