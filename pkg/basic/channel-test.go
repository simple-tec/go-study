package basic

import "fmt"

func Chan_test1() {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		c1 <- i
		c2 <- i
	}
	for {
		select {
		case <-c1:
			fmt.Println("random 1")
		case <-c2:
			fmt.Println("random 2")
		default:
			//fmt.Println("default")
		}
	}
}
