package basic

func Test_if() {
	a, b := false, true
	if a && b != true {
		//if a && (b != true) {
		println("debug 1, true")
		return
	}
	println("debug 2, false")
}
