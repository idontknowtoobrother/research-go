package main

import (
	"fmt"
)

func main() {

	ch := make(chan int, 2)

	go func() {
		ch <- 123561
	}()

	fmt.Println("waiting for go channel.")
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
