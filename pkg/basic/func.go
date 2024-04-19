package basic

import (
	"fmt"
	"reflect"
)

/*
 */
func Add(a int, b int) (int, error) {
	return a + b, nil
}

// AddWrap()的返回值是函数
func AddWrap() func(c int, d int) (int, error) {
	fmt.Printf("AddWrap\n")
	return func(a int, b int) (int, error) {
		fmt.Printf("add run\n")
		return a + b, nil
	}
}

func AddWrap2(a int, b int, addFunc func(c int, d int) (int, error)) (int, error) {
	return addFunc(a, b)
}

func Test_func() {
	// 函数正常用法
	a := 1
	b := 2
	c, _ := Add(a, b)
	fmt.Printf("c=%d\n", c)

	// 函数存在变量里
	addFunc := func(a int, b int) (int, error) {
		return Add(a, b)
	}

	d, _ := addFunc(a, b)
	fmt.Printf("d=%d\n", d)

	// 函数作为函数的返回值
	addFunc = AddWrap()
	e, _ := addFunc(a, b)
	fmt.Printf("e=%d\n", e)

	// 函数作为函数的参数
	f, _ := AddWrap2(a, b, Add)
	fmt.Printf("f=%d\n", f)

	// 函数类型判断
	funcA := func(a int, b int) (int, error) {
		return 0, nil
	}
	funcB := func(a int, b int) (int, error) {
		return 1, nil
	}
	if reflect.TypeOf(funcA) == reflect.TypeOf(funcB) {
		fmt.Printf("funcA equals to funcB\n")
	} else {
		fmt.Printf("funcA not equals to funcB\n")
	}

}
