package basic

import (
	"errors"
	"fmt"
)

func Test_error() {
	err1 := errors.New("your first demo error")
	fmt.Printf("%s\n", err1.Error())
}
