package convey

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {
	convey.
		Convey("将两数相加", t, func() {
			convey.So(Add(1, 2), convey.ShouldEqual, 3)
		})
}

func TestSubtract(t *testing.T) {
	convey.
		Convey("将两数相减", t, func() {
			convey.So(Subtract(1, 2), convey.ShouldEqual, -1)
		})
}

func TestMultiply(t *testing.T) {
	convey.
		Convey("将两数相乘", t, func() {
			convey.So(Multiply(3, 2), convey.ShouldEqual, 6)
		})
}

func TestDivision(t *testing.T) {
	convey.
		Convey("将两数相除", t, func() {
			convey.
				Convey("除以非 0 数", func() {
					num, err := Division(10, 2)
					convey.So(err, convey.ShouldBeNil)
					convey.So(num, convey.ShouldEqual, 5)
				})
			convey.
				Convey("除以 0", func() {
					_, err := Division(10, 0)
					convey.So(err, convey.ShouldNotBeNil)
				})
		})
}
