package strutils

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRand(t *testing.T) {
	Convey("When two random strings are generated", t, func() {
		s1 := Rand(10)
		s2 := Rand(10)

		Convey("They should have length 10", func() {
			So(s1, ShouldHaveLength, 10)
			So(s2, ShouldHaveLength, 10)
		})

		Convey("They should not be the same", func() {
			So(s1, ShouldNotEqual, s2)
		})
	})

	Convey("When too many strings are generated", t, func() {
		start := time.Now()
		for i := 0; i < 4500; i++ {
			Rand(10)
		}
		end := time.Now()

		Convey("The operation should not take more than 10 miliseconds", func() {
			So(start, ShouldHappenWithin, 10*time.Millisecond, end)
		})
	})
}
