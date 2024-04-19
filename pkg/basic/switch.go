package basic

/*
c语言：switch求值结果只能是int或枚举类型；
go语言：只要类型支持比较操作就可以；
c语言：如果匹配到的case对应的代码分支中没有显式调用break语句，那么代码将继续执行下一个case的代码分支；
go语言：每个case对应的分支代码执行完后就结束switch语句，无需显式调用break语句。
如果需要执行下一个case的代码逻辑，你可以显式使用Go提供的关键字fallthrough来实现
*/

func checkWorkday(a int) {
	switch a {
	case 1, 2, 3, 4, 5: //表达式列表
		println("work day")
		fallthrough
		// 由于fallthrough的存在，Go会直接执行case 6,7对应的代码分支
	case 6, 7:
		println("rest day")
	default:
		println("invalid day")
	}
}

func Test_switch() {
	checkWorkday(1)
}
