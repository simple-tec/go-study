package basic

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
string类型的变量其实是一个标识，它本身并不真正存储字符串数据，而仅是由一个指向底层存储的指针和字符串的长度字段组成的。
底层实现类型是StringHeader
*/

func Test_string() {
	var s1 string = "hello"
	sHeader := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	fmt.Printf("len: %d\n", sHeader.Len)

	//s1[0] = 'k'  // 错误：字符串的内容是不可改变的
	//字符迭代
	for i, v := range s1 {
		fmt.Printf("index: %d, value: 0x%x\n", i, v)
	}
	//字符串比较
	if s1 == "hello" {
		fmt.Printf("equals\n")
	}
	//字符串连接
	s2 := s1 + " world"
	fmt.Printf("s2: %s\n", s2)
}
