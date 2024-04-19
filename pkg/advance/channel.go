package advance

import (
	"fmt"
	"time"
)

func Test_channel() {
	ch1 := make(chan int) // 无缓冲channel
	// ch2 := make(chan int, 5) //带缓冲channel
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 13 // 将发送操作放入一个新goroutine中执行
	}()
	fmt.Println("begin recv")
	n := <-ch1
	fmt.Printf("recv from channel: %d\n", n)
}

func Test_channel1() {
	// ch1 := make(chan<- int, 1) // 只发送channel类型
	// ch2 := make(<-chan int, 1) // 只接收channel类型

	// <-ch1     // invalid operation: <-ch1 (receive from send-only type chan<- int)
	// ch2 <- 13 // invalid operation: ch2 <- 13 (send to receive-only type <-chan int)
}
