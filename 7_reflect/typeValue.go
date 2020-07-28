package __reflect

import (
	"fmt"
	"reflect"
)

func GetType() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)
	// 	=== RUN   TestGetType
	// int
	// int
}

func GetValue() {
	t := reflect.ValueOf(3)
	fmt.Println(t)
	fmt.Println(t.String())
	// 	=== RUN   TestGetType
	// int
	// int
}

func GetMethod(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()

	fmt.Printf("type: %v\n", t)

	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("func: %s %s %s\n", t, t.Method(i).Name, methodType.String())
	}
}

// type: time.Duration
// func: time.Duration Hours func() float64
// func: time.Duration Microseconds func() int64
// func: time.Duration Milliseconds func() int64
// func: time.Duration Minutes func() float64
// func: time.Duration Nanoseconds func() int64
// func: time.Duration Round func(time.Duration) time.Duration
// func: time.Duration Seconds func() float64
// func: time.Duration String func() string
// func: time.Duration Truncate func(time.Duration) time.Duration

func SetValue() {
	x := 2
	a := reflect.ValueOf(2)
	b := reflect.ValueOf(x)
	c := reflect.ValueOf(&x)
	d := c.Elem()

	fmt.Println(a, a.CanAddr())
	fmt.Println(b, b.CanAddr())
	fmt.Println(c, c.CanAddr())
	fmt.Println(d, d.CanAddr())

	d.Set(reflect.ValueOf(3))
	fmt.Println(x)
	d.SetInt(9)
	fmt.Println(x)
}
