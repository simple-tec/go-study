package advance

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func Test_interface() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()
	phone = new(IPhone)
	phone.call()

	// type assert 1
	v1, ok1 := phone.(*NokiaPhone)
	fmt.Printf("v1 type: %T, value: %#v, ok1=%t\n", v1, v1, ok1)
	v2, ok2 := phone.(*IPhone)
	fmt.Printf("v2 type: %T, value: %#v, ok2=%t\n", v2, v2, ok2)

	// type asset 2
	v3, ok3 := phone.(interface{})
	fmt.Printf("v3 type: %T, value: %#v, ok3=%t\n", v3, v3, ok3)

	type Phone2 interface {
		call2()
	}
	v4, ok4 := phone.(Phone2)
	fmt.Printf("v4 type: %T, value: %#v, ok4=%t\n", v4, v4, ok4)
}
