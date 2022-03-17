package main

import "fmt"

type C[T any] chan T

func main() {
	c1 := make(C[int], 10)
	c1 <- 1
	c1 <- 2

	c2 := make(C[string], 10)
	c2 <- "hello"
	c2 <- "world"

	fmt.Println(<-c1, <-c2)
	fmt.Println(<-c1, <-c2)
}
