package pkg

func fn1() {
	return // MATCH /redundant return/
}

func fn2(a int) {
	return // MATCH /redundant return/
}

func fn3() int {
	return 3
}

func fn4() (n int) {
	return
}

func fn5(b bool) {
	if b {
		return
	}
}

func fn6() {
	return
	println("foo")
}

func fn7() {
	return
	println("foo")
	return // MATCH /redundant return/
}

func fn8() {
	_ = func() {
		return // MATCH /redundant return/
	}
}
