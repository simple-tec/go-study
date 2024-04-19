package advance

import (
	"errors"
	"fmt"
)

type MyError struct {
	error
}

var Error1 = MyError{
	error: errors.New("error1"),
}

func returnErr() error {
	var p *MyError = nil
	fmt.Printf("p type: %T, value: %+v\n", p, p)
	return p
}
func Test_nil() {
	err := returnErr()
	fmt.Printf("error type: %T, value: %+v\n", err, err)
	// p为nil, returnErr返回p,为什么这里的err != nil呢？
	if err != nil {
		fmt.Printf("error occur: %+v\n", err)
		return
	}
	fmt.Println("ok")
}

func Test_nil1() {
	var i interface{}
	var err error
	var perr *error
	fmt.Printf("i type: %T, value: %+v\n", i, i)
	fmt.Printf("err type: %T, value: %+v\n", err, err)
	fmt.Printf("perr type: %T, value: %+v\n", perr, perr)
	println("i = err:", i == err) // true

	var if1 interface{}
	var if2 interface{}
	var n, m int = 1, 2
	if1 = n
	if2 = m
	println("if1 = if2:", if1 == if2) // false
	if2 = 1
	println("if1 = if2:", if1 == if2) // true
	if2 = int64(1)
	println("if1 = if2:", if1 == if2) // false

}
