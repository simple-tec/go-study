package basic

import (
	"fmt"
	"unsafe"
)

type Author struct {
	Name string
	Age  int
}
type Book struct {
	Book_id int
	Title   string
	Author  Author
}

func Test_struct() {
	// 空结构体的例子
	type Empty struct{}
	var s Empty
	fmt.Println(unsafe.Sizeof(s)) // 0

	// 结构体变量定义和初始化
	author := Author{
		Name: "john",
		Age:  35,
	}
	book := Book{
		Book_id: 2001,
		Title:   "go语言编程",
		Author:  author,
	}
	fmt.Printf("book=%#v\n", book)

	//
	fmt.Println(unsafe.Sizeof(book))
	fmt.Println(unsafe.Offsetof(book.Book_id))
	fmt.Println(unsafe.Offsetof(book.Title))
	fmt.Println(unsafe.Offsetof(book.Author))
}
