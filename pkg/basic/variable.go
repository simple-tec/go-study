package basic

import "fmt"

var (
	default404Body = []byte("404 page not found")
	default405Body = []byte("405 method not allowed")
)

func Test_variable() {
	var a1 int = 10    // 完整格式
	var a2 int         // int零值为0
	var a3 = 10        // 省略类型
	var a4 = int32(10) // 指定类型
	a5 := 10
	fmt.Printf("a1 Value: %d, Type: %T\n", a1, a1)
	fmt.Printf("a2 Value: %d, Type: %T\n", a2, a2)
	fmt.Printf("a3 Value: %d, Type: %T\n", a3, a3)
	fmt.Printf("a4 Value: %d, Type: %T\n", a4, a4)
	fmt.Printf("a5 Value: %d, Type: %T\n", a5, a5)
}
