package strutils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStringUtils(t *testing.T) {
	Convey("Given a big number", t, func() {
		number := "10000000"
		Convey("When the number is grouped into 3 digits", func() {
			result := GroupDigits(number, ",", 3)

			Convey("The result should be 10,000,000", func() {
				So(result, ShouldEqual, "10,000,000")
			})
		})
		Convey("When the number is grouped into 2 digits", func() {
			result := GroupDigits(number, ",", 2)

			Convey("The result should be 10,00,00,00", func() {
				So(result, ShouldEqual, "10,00,00,00")
			})
		})
		Convey("When the number is grouped into 4 digits", func() {
			result := GroupDigits(number, ",", 4)

			Convey("The result should be 1000,0000", func() {
				So(result, ShouldEqual, "1000,0000")
			})
		})
		Convey("When the number is grouped into 5 digits", func() {
			result := GroupDigits(number, ",", 5)

			Convey("The result should be 100,00000", func() {
				So(result, ShouldEqual, "100,00000")
			})
		})
	})
}

func TestUnderScore(t *testing.T) {
	Convey("Underscore", t, func() {
		So("i_love_golang_and_json_so_much", ShouldEqual, UnderScore("ILoveGolangAndJSONSoMuch"))
		So("i_love_json", ShouldEqual, UnderScore("ILoveJSON"))
		So("json", ShouldEqual, UnderScore("json"))
		So("json", ShouldEqual, UnderScore("JSON"))
		So("привет_мир", ShouldEqual, UnderScore("ПриветМир"))
	})
}
