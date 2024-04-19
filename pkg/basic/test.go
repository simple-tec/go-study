package basic

import "fmt"

func test1(array [5]int) {
	array[0] = 0
}
func test2() bool {
	// ..
	return true
}
func Test() {
	a := 1
	switch a {
	case 1, 2, 3, 4, 5:
		fmt.Println("this is work day")

	case 6, 7:
		fmt.Println("this is rest day")
	}
	/*
		var arr1 = [5]int{1, 2, 3, 4, 5}
		test1(arr1)
		var i int
		for i = 0; i < 5; i++ {
			fmt.Printf("arr1[%d]=%d\n", i, arr1[i])
		}
	*/
	/*
		var err error
		err = errors.New("a error")
		fmt.Printf("%T\n", err) // *errors.errorString
	*/
}
