package basic

import "fmt"

func Test_for() {
	var sum int
	for i := 0; i < 10; i++ {
		sum += i
	}
	println(sum)

	// for range, slice
	var slice = []int{1, 2, 3, 4, 5}
	for i, v := range slice {
		slice[i] = i + 1
		fmt.Printf("index %d = %d\n", i, v)
	}

	for i, v := range slice {
		fmt.Printf("index %d = %d\n", i, v)
	}

	var i int
	var v int
	for i, v = range slice {
		fmt.Printf("index %d = %d\n", i, v)
	}

	// for range, string
	var str1 = "中国的首都是北京"
	for i, v := range str1 {
		fmt.Printf("index %d = %s\n", i, string(v))
	}

	// for range, map
	var map1 = map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	for k, v := range map1 {
		fmt.Printf("%s = %s\n", k, v)
	}

}
