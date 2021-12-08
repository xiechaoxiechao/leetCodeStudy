package main

import (
	"fmt"
	"time"
)

func main() {

	var c=make(chan int)
	go func() {
		for{
			select {
			case <-c:
				fmt.Println("get")
			}
		}

	}()
	for i:=1;i<10;i++{
		c<-0
		time.Sleep(2*time.Second)
	}
	time.Sleep(time.Minute)

}
