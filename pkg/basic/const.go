package basic

import "fmt"

func Test_const() {
	const c1 int = 13
	const c2 = 13 // 编译器会有隐式转换
	fmt.Printf("c1 Value: %d, Type: %T\n", c1, c1)
	fmt.Printf("c2 Value: %d, Type: %T\n", c2, c2)
}
