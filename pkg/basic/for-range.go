package basic

import (
	"fmt"
	"time"
)

func Test_forrange() {
	var m = []int{1, 2, 3, 4, 5}

	// i,v 在for循环里一次定义，多次重用
	for i, v := range m {
		/*
			goroutine执行的闭包函数引用了它的外层(也就是for循环)的变量i、v
			这样的话，变量i、v在主Goroutine和新启动的Goroutine之间实现了共享
			而i, v 值在整个循环过程中是重用的，仅有一份。
			在for range循环结束后，i = 4, v = 5，
			因此各个 Goroutine 在等待3秒后进行输出的时候，输出的是 i, v 的最终值。
		*/
		go func() {
			time.Sleep(time.Second * 5)
			fmt.Println(i, v)
		}()
	}

	time.Sleep(time.Second * 20)
}
