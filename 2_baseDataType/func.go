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
