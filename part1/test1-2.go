package main

import "fmt"

func sum(values []int, res chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}

	res <- sum
}

func main() {
	var values [5]int = [5]int{1, 2, 3, 4, 5}

	res := make(chan int, 2)

	go sum(values[:len(values)/2], res)
	go sum(values[len(values)/2:], res)

	sum1, sum2 := <-res, <-res
	fmt.Println(sum1, sum2, sum1+sum2)
}
