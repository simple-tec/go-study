package basic

import "fmt"

func Test_slice_append() {
	slice2 := []byte{5, 6, 7}
	slice1 := make([]byte, 2, 4)
	fmt.Printf("%p,%v\n", slice1, slice1)
	slice1 = append(slice1, 1)
	fmt.Printf("%p,%v\n", slice1, slice1)
	slice1 = append(slice1, 2)
	fmt.Printf("%p,%v\n", slice1, slice1)
	slice1 = append(slice1, 3)
	fmt.Printf("%p,%v\n", slice1, slice1)
	slice1 = append(slice1, slice2...)
	fmt.Printf("%p,%v\n", slice1, slice1)
}

func Test_slice() {
	var slice = []int{1, 2, 3, 4, 5}
	change_slice(slice)
	fmt.Println(slice)
}

func change_slice(arr []int) {
	arr[0] = 11
}
