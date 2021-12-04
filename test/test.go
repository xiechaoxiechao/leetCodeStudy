package main

import "fmt"

func main() {
	a := make(chan int, 2)
	a <- 3
	close(a)
	for b := range a {
		fmt.Println(b)
	}

}
