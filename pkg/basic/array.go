package basic

import "fmt"

func Test_array() {
	var arr = [5]int{1, 2, 3, 4, 5}
	change_array(arr)
	fmt.Println(arr)
}

func change_array(arr [5]int) {
	arr[0] = 11
}
