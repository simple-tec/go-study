package basic

/*
不带label的break语句跳出的是break语句所在的最内层的for、switch或select
*/
func Test_nolabel_break() {
	var slice1 = []int{1, 13, 6, 3, 5, 18}
	var result int

	// 找到整数切片中第一个偶数
	for i := 0; i < len(slice1); i++ {
		switch slice1[i] % 2 {
		case 0:
			result = slice1[i]
			break
		case 1:
			// do nothing
		}
	}
	println(result)
}

/*
带label的break语句才能跳出外层的for、switch或select
*/
func Test_label_break() {
	var slice1 = []int{1, 13, 6, 3, 5, 18}
	var result int

	// 找到整数切片中第一个偶数
	// label: loop在for循环的外面，指代for循环的执行
loop:
	for i := 0; i < len(slice1); i++ {
		switch slice1[i] % 2 {
		case 0:
			result = slice1[i]
			// 执行break loop，将停止for循环的执行
			break loop
		case 1:
			// do nothing
		}
	}
	println(result)
}
