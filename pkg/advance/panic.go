package advance

import (
	"time"
)

func call1() {
	println("call call1")
	call2()
	println("exit call1")
}

func call2() {
	// 通过recover()函数捕捉panic并恢复程序正常执行
	/*
		defer func() {
			if e := recover(); e != nil {
				fmt.Println("recover the panic:", e)
			}
		}()
	*/
	println("call call2")
	panic("panic occurs in call2")
	call3()
	println("exit call2")
}

func call3() {
	println("call call3")
	println("exit call3")
}

func Test_panic() {
	println("call test panic")
	call1()
	println("exit test panic")
}

func Test_panic1() {
	/* Test_panic1向我们展示了：
	即使是在某个Goroutine中发生未被恢复的panic，整个程序都将崩溃退出
	*/
	println("call test panic")
	go call1()
	time.Sleep(time.Second)
	println("exit test panic")
}
