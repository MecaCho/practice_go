package __baseDataType

func GetFuncs() []func() {

	var funs []func()
	for i := 0; i < 3; i++ {
		funs = append(funs, func() {
			println(&i, i)
		})
	}
	return funs

}

func RunFuncs() {
	funs := GetFuncs()
	for _, f := range funs {
		f()
	}
}

func TwoFunc(x int) (func(), func()) {
	return func() {
			println(x)
			x += 10
		}, func() {
			println(x)
		}
}

func RunTwoFunc() {

	a, b := TwoFunc(100)
	a()
	b()

}

// 匿名函数的调用 匿名函数是不能自己调用自己的, f(i - 1)其实是上面var定义的 f
var f = func(i int) {
	print("x")
}

func FuncInFunc() {

	f := func(i int) {
		print(i)
		if i > 0 {
			f(i - 1)
		}
	}
	f(10)

}
