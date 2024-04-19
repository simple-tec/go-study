package advance

import "fmt"

/*
defer要点：
1、多个defer函数，按后进先出的顺序被程度调度执行
2、defer关键字后面的表达式，是在将deferred函数注册到deferred函数栈的时候进行求值的
*/

func test1() {
	// 每当 defer 将 fmt.Println 注册到 deferred 函数栈的时候
	// 都会对fmt.Println(i)里的i进行求值
	for i := 0; i <= 3; i++ {
		defer fmt.Println(i)
	}
	/*
		defer fmt.Println(0)
		defer fmt.Println(1)
		defer fmt.Println(2)
		defer fmt.Println(3)
	*/
}

func test2() {
	// 每当 defer 将 fmt.Println 注册到 deferred 函数栈的时候
	// 都会对func的参数i进行求值
	for i := 0; i <= 3; i++ {
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}
	/*
		func(0)
		func(1)
		func(2)
		func(3)
	*/
}

func test3() {
	// 每当 defer 将 fmt.Println 注册到 deferred 函数栈的时候
	// 这里的func()没有参数，所以不求值
	// 最后即将退出test()的时候执行fmt.Println(i)，此时i等于4
	for i := 0; i <= 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
	// i = 4
	/*
		func()
		func()
		func()
		func()
	*/
}

func Test_defer() {
	//fmt.Println("test1 result:")
	//test1()
	//fmt.Println("\ntest2 result:")
	//test2()
	fmt.Println("\ntest3 result:")
	test3()
}
